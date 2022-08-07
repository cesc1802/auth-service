package usecase

import (
	"context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/Permission/domain"
	"github.com/cesc1802/auth-service/features/v1/Permission/dto"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UpdateStore interface {
	generic.IFindOneByConditionStore[domain.Permission]
	generic.IUpdateStore[domain.Permission]
}

type ucUpdatePermission struct {
	store UpdateStore
}

func NewUseCaseUpdatePermission(store UpdateStore) *ucUpdatePermission {
	return &ucUpdatePermission{
		store: store,
	}
}

func (uc *ucUpdatePermission) UpdatePermission(ctx context.Context, id uint, form *dto.UpdatePermissionRequest) error {
	Permission, err := uc.store.FindOneByCondition(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})

	if err != nil {
		return common.ErrCannotGetEntity(domain.EntityName, err)
	}

	if err := copier.Copy(Permission, form); err != nil {
		return common.ErrCopyData
	}
	if err := uc.store.Update(ctx, Permission); err != nil {
		return common.ErrCannotUpdateEntity(domain.EntityName, err)
	}

	return nil
}
