package storage

import (
	"github.com/cesc1802/auth-service/features/v1/role_permissions/domain"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"gorm.io/gorm"
)

type mySqlRolePermissionStore struct {
	*generic.CRUDStore[domain.RolePermission]
}

func NewMySqlRolePermissionStore(db *gorm.DB) *mySqlRolePermissionStore {
	return &mySqlRolePermissionStore{
		generic.NewCRUDStore[domain.RolePermission](db),
	}
}
