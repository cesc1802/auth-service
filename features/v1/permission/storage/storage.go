package storage

import (
	"github.com/cesc1802/auth-service/features/v1/Permission/domain"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"gorm.io/gorm"
)

type mySqlPermissionStore struct {
	*generic.CRUDStore[domain.Permission]
}

func NewMySqlPermissionStore(db *gorm.DB) *mySqlPermissionStore {
	return &mySqlPermissionStore{
		generic.NewCRUDStore[domain.Permission](db),
	}
}
