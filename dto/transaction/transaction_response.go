package transactiondto

import "dumbmerch/models"

type TransactionResponse struct {
	Id         int    `json:"id_trans" form:"id_trans" gorm:"primary_key:auto_increment"`
	CounterQty int    `json:"counter_qty" form:"counter_qty"`
	Total      int    `json:"total" form:"total"`
	Status     string `json:"status" form:"status"`
	Attachment string `json:"attachment" form:"attachment"`
	IdTrip     int    `json:"id_trip" form:"id_trip"`
	Trip       models.TripResponse
	IdUser     int `json:"id_user" form:"id_user"`
	User       models.UserResponse
}
