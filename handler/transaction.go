package handler

import (
	resultdto "dumbmerch/dto/result"
	transactiondto "dumbmerch/dto/transaction"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gopkg.in/gomail.v2"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type HandlerTransactions struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *HandlerTransactions {
	return &HandlerTransactions{TransactionRepository}
}

func (h *HandlerTransactions) FindTransaction(c echo.Context) error {
	transaction, err := h.TransactionRepository.FindTransaction()

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusOK,
			Message: "Waduh"})
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: transaction})

}
func (h *HandlerTransactions) GetTransByUser(c echo.Context) error {
	// id, _ := strconv.Atoi(c.Param("id"))
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)
	transaction, err := h.TransactionRepository.GetTransByUser(int(userId))

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusOK,
			Message: "Waduh"})
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: transaction})

}

func (h *HandlerTransactions) FindTransactionId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, _ := h.TransactionRepository.FindTransactionId(id)

	if transaction.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Data Gaada Bos"})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: transaction})
}

func (h *HandlerTransactions) DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, _ := h.TransactionRepository.FindTransactionId(id)

	if transaction.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "KOSONG OM"})
	}

	data, err := h.TransactionRepository.DeleteTransaction(id, transaction)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)})
}

func (h *HandlerTransactions) CreateTransaction(c echo.Context) error {
	request := new(transactiondto.CreateTransaction)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	request.IdUser = int(userId)
	request.Status = "pending" // next time we will cover this in payment gateway material

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	trips, err := h.TransactionRepository.GetTripId(request.IdTrip)

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}
	users, err := h.TransactionRepository.GetUserId(int(userId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	var transactionIsMatch = false
	var transactionId int
	for !transactionIsMatch {
		transactionId = int(time.Now().Unix())
		transactionData, _ := h.TransactionRepository.FindTransactionId(transactionId)
		if transactionData.Id == 0 {
			transactionIsMatch = true
		}
	}

	request.Status = "pending"

	fmt.Println(userId)
	Transaction := models.Transaction{
		Id:         transactionId,
		User:       users,
		CounterQty: request.CounterQty,
		Total:      request.Total,
		Status:     request.Status,
		Attachment: request.Attachment,
		IdTrip:     request.IdTrip,
		Trip:       trips,
		IdUser:     int(userId),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	data, err := h.TransactionRepository.CreateTransaction(Transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}
	fmt.Println(data, "DATA INI CUKK")
	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(data.Id),
			GrossAmt: int64(data.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: data.User.Name,
			Email: data.User.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: snapResp})

	// return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)})
}

func (h *HandlerTransactions) Notification(c echo.Context) error {
	var notificationPayload map[string]interface{}

	if err := c.Bind(&notificationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	order_id, _ := strconv.Atoi(orderId)
	transaction, _ := h.TransactionRepository.FindTransactionId(order_id)

	fmt.Print("ini payloadnya", notificationPayload)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			// TODO set transaction status on your database to 'challenge'
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			h.TransactionRepository.UpdateTransaction("pending", order_id)
		} else if fraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			SendMail("success", transaction)
			h.TransactionRepository.UpdateTransaction("success", order_id)
		}
	} else if transactionStatus == "settlement" {
		// TODO set transaction status on your databaase to 'success'
		SendMail("success", transaction)
		h.TransactionRepository.UpdateTransaction("success", order_id)
	} else if transactionStatus == "deny" {
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
		h.TransactionRepository.UpdateTransaction("failed", order_id)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		// TODO set transaction status on your databaase to 'failure'
		h.TransactionRepository.UpdateTransaction("failed", order_id)
	} else if transactionStatus == "pending" {
		// TODO set transaction status on your databaase to 'pending' / waiting payment
		h.TransactionRepository.UpdateTransaction("pending", order_id)
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: notificationPayload})
}

func SendMail(status string, transaction models.Transaction) {

	if status != transaction.Status && (status == "success") {
		var CONFIG_SMTP_HOST = "smtp.gmail.com"
		var CONFIG_SMTP_PORT = 587
		var CONFIG_SENDER_NAME = "DumbMerch <demo.dumbways@gmail.com>"
		var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
		var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")

		var title = transaction.Trip.Title
		var total = strconv.Itoa(transaction.Total)

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", "rodyulo20@gmail.com")
		mailer.SetHeader("Subject", "Transaction Status")
		mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
	  <html lang="en">
		<head>
		<meta charset="UTF-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Document</title>
		<style>
		  h1 {
		  color: brown;
		  }
		</style>
		</head>
		<body>
		<h2>Product payment :</h2>
		<ul style="list-style-type:none;">
		  <li>Name : %s</li>
		  <li>Total payment: Rp.%s</li>
		  <li>Status : <b>%s</b></li>
		</ul>
		</body>
	  </html>`, title, total, status))

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent! to " + "rodyulo20@gmail.com")
	}
}

func convertResponseTransaction(Transaction models.Transaction) transactiondto.TransactionResponse {
	return transactiondto.TransactionResponse{
		Id:         Transaction.Id,
		CounterQty: Transaction.CounterQty,
		Total:      Transaction.Total,
		Status:     Transaction.Status,
		Attachment: Transaction.Attachment,
		IdTrip:     Transaction.IdTrip,
		Trip:       Transaction.Trip,
		IdUser:     Transaction.IdUser,
		User:       Transaction.User,
	}
}
