package entities

import "github.com/cesc1802/auth-service/common"

type Permission struct {
	common.BaseModel
	Name        string  `gorm:"column:name"`
	Description *string `gorm:"column:description"`
}

func (r Permission) TableName() string {
	return "permissions"
}

func (r Permission) EntityName() string {
	return "permission"
}
