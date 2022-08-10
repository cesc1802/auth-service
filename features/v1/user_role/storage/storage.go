package storage

import (
	"github.com/cesc1802/auth-service/features/v1/user_role/domain"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"gorm.io/gorm"
)

type mySqlUserRoleStore struct {
	*generic.CRUDStore[domain.UserRole]
}

func NewMySqlUserRoleStore(db *gorm.DB) *mySqlUserRoleStore {
	return &mySqlUserRoleStore{
		generic.NewCRUDStore[domain.UserRole](db),
	}
}
