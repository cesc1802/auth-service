package entities

import (
	"github.com/cesc1802/auth-service/pkg/database"
	"time"
)

type RolePermission struct {
	RoleID       uint               `json:"role_id" gorm:"primaryKey,column:role_id"`
	PermissionID uint               `json:"permission_id" gorm:"primaryKey,column:permission_id"`
	Status       int                `json:"status" gorm:"column:status;default:1;"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt    *time.Time         `json:"updated_at" gorm:"column:updated_at;"`
	DeletedAt    database.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

func (rp RolePermission) TableName() string {
	return "role_permissions"
}

func (rp RolePermission) EntityName() string {
	return "role_permission"
}
