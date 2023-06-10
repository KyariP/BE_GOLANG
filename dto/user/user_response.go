package userdto

import "dumbmerch/models"

type UserResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	// TransactionId int                `json:"transaction_id"`
	Transaction []models.Transaction `json:"transaction"`
}
