package models

type Role struct {
	Id   int    `json:"role_id" form:"role_id" gorm:"primary_key:auto_increment"`
	Name string `json:"role_name" form:"role_name" gorm:"type: varchar(255)"`
}

type RoleResponse struct {
	Id   int    `json:"role_id"`
	Name string `json:"role_name"`
}

func (RoleResponse) TableName() string {
	return "roles"
}
