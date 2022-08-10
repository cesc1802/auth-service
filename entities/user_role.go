package entities

import (
	"time"

	"github.com/cesc1802/auth-service/pkg/database"
)

type UserRole struct {
	RoleID    uint               `json:"role_id" gorm:"column:role_id"`
	UserID    uint               `json:"user_id" gorm:"column:user_id"`
	Status    int                `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time         `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time         `json:"updated_at" gorm:"column:updated_at;"`
	DeletedAt database.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

func (r UserRole) TableName() string {
	return "user_roles"
}

func (r UserRole) EntityName() string {
	return "user_role"
}
