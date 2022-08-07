package usecase

import (
	"context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/Permission/domain"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"gorm.io/gorm"
)

type DeletePermissionStore interface {
	generic.IFindOneByConditionStore[domain.Permission]
	generic.IDeleteStore[domain.Permission]
}

type ucDeletePermission struct {
	store DeletePermissionStore
}

func NewUseCaseDeleteStore(store DeletePermissionStore) *ucDeletePermission {
	return &ucDeletePermission{
		store: store,
	}
}

func (uc *ucDeletePermission) DeletePermission(ctx context.Context, id uint) error {
	_, err := uc.store.FindOneByCondition(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})

	if err != nil && err != common.ErrRecordNotFound {
		return common.ErrCannotGetEntity(domain.EntityName, err)
	}

	if err := uc.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(domain.EntityName, err)
	}

	return nil
}
