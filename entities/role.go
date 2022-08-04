package entities

import "github.com/cesc1802/auth-service/common"

type Role struct {
	common.BaseModel
	Name        string  `gorm:"column:name"`
	Description *string `gorm:"column:description"`
}

func (r Role) TableName() string {
	return "roles"
}

func (r Role) EntityName() string {
	return "role"
}
