package models

import "time"

type Transaction struct {
	Id         int          `json:"id_trans" form:"id_trans" gorm:"primary_key:auto_increment"`
	IdUser     int          `json:"id_user" form:"id_user"`
	User       UserResponse `json:"user" form:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:IdUser"`
	CounterQty int          `json:"counter_qty" form:"counter_qty"`
	Total      int          `json:"total" form:"total"`
	Status     string       `json:"status" form:"status"`
	Attachment string       `json:"attachment" form:"attachment"`
	IdTrip     int          `json:"id_trip" form:"id_trip"`
	Trip       TripResponse `json:"trip" form:"trip" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:IdTrip"`
	CreatedAt  time.Time    `json:"-"`
	UpdatedAt  time.Time    `json:"-"`
}

type TransactionResponse struct {
	Id         int          `json:"id_trans"`
	IdUser     int          `json:"id_user"`
	User       UserResponse `json:"user"`
	CounterQty int          `json:"counter_qty"`
	Total      int          `json:"total"`
	Status     string       `json:"status"`
	Attachment string       `json:"attachment"`
	IdTrip     int          `json:"id_trip"`
	Trip       TripResponse `json:"trip"`
}

func (TransactionResponse) TableName() string {
	return "trans"
}
