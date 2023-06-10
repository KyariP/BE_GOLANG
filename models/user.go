package models

import (
	"time"
)

type User struct {
	Id          int           `json:"id" form:"id" gorm:"primary_key:auto_increment"`
	Name        string        `json:"name" form:"name" gorm:"type: varchar(255)"`
	Email       string        `json:"email" form:"email" gorm:"type: varchar(255)"`
	Password    string        `json:"password" form:"password" gorm:"type: varchar(255)"`
	Phone       string        `json:"phone" form:"phone" gorm:"type: varchar(255)"`
	Address     string        `json:"address" form:"address" gorm:"type: varchar(255)"`
	RoleId      int           `json:"role_id" form:"role_id"`
	RoleName    RoleResponse  `json:"role_name" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:RoleId"`
	Transaction []Transaction `json:"transaction" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:IdUser"`
	CreatedAt   time.Time     `json:"-"`
	UpdatedAt   time.Time     `json:"-"`
}

type UserResponse struct {
	Id       int          `json:"id" form:"id"`
	Name     string       `json:"name"`
	Email    string       `json:"email"`
	Password string       `json:"password"`
	Phone    string       `json:"phone"`
	Address  string       `json:"address"`
	RoleId   int          `json:"role_id" form:"role_id"`
	RoleName RoleResponse `json:"role_name" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:RoleId"`
}

func (UserResponse) TableName() string {
	return "users"
}
