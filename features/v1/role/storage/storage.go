package storage

import (
	"github.com/cesc1802/auth-service/features/v1/role/domain"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"gorm.io/gorm"
)

type mySqlRoleStore struct {
	*generic.CRUDStore[domain.Role]
}

func NewMySqlRoleStore(db *gorm.DB) *mySqlRoleStore {
	return &mySqlRoleStore{
		generic.NewCRUDStore[domain.Role](db),
	}
}
