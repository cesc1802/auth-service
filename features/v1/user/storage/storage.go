package storage

import (
	"github.com/cesc1802/auth-service/features/v1/user/domain"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"gorm.io/gorm"
)

type mySqlUserStore struct {
	*generic.CRUDStore[domain.User]
}

func NewMySqlUserStore(db *gorm.DB) *mySqlUserStore {
	return &mySqlUserStore{
		generic.NewCRUDStore[domain.User](db),
	}
}
