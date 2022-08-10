package usecase

import (
	"context"

	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/user_role/domain"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"gorm.io/gorm"
)

type GetRolesStore interface {
	generic.IFindAllStore[domain.UserRole]
}

type ucGetRoles struct {
	store GetRolesStore
}

func NewUseCaseGetRoles(store GetRolesStore) *ucGetRoles {
	return &ucGetRoles{
		store: store,
	}
}

func (uc *ucGetRoles) GetRoles(ctx context.Context, userID uint) ([]domain.UserRole, error) {
	roles, _, err := uc.store.FindAll(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userID)
	})
	if err != nil {
		return nil, common.ErrCannotGetEntity(domain.EntityName, err)
	}

	return roles, nil
}
