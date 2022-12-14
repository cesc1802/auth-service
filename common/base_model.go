package common

import (
	"github.com/cesc1802/auth-service/pkg/database"
	"time"
)

type BaseModel struct {
	ID        uint               `json:"id" gorm:"primaryKey"`
	Status    int                `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time         `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time         `json:"updated_at" gorm:"column:updated_at;"`
	DeletedAt database.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

type BaseModelCreate struct {
}
