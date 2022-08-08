package entities

import "github.com/cesc1802/auth-service/common"

type UserRole struct {
	common.BaseModel
	RoleID int `gorm:"column:role_id"`
	UserID int `gorm:"column:user_id"`
}

func (r UserRole) TableName() string {
	return "user_roles"
}

func (r UserRole) EntityName() string {
	return "user_role"
}
