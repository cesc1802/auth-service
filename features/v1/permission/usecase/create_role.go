package usecase

import (
	"context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/entities"
	"github.com/cesc1802/auth-service/features/v1/Permission/domain"
	"github.com/cesc1802/auth-service/features/v1/Permission/dto"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"gorm.io/gorm"
)

type CreatePermissionStore interface {
	generic.IFindOneByConditionStore[domain.Permission]
	generic.ICreateStore[domain.Permission]
}

type ucCreatePermission struct {
	store CreatePermissionStore
}

func NewUseCaseCreatePermission(store CreatePermissionStore) *ucCreatePermission {
	return &ucCreatePermission{
		store: store,
	}
}

func (uc *ucCreatePermission) CreatePermission(ctx context.Context, form *dto.CreatePermissionRequest) error {
	Permission, err := uc.store.FindOneByCondition(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("name = ?", form.Name)
	})

	if err != nil && err != common.ErrRecordNotFound {
		return common.ErrCannotGetEntity(domain.EntityName, err)
	}

	if Permission != nil {
		return domain.ErrPermissionNameIsExisting
	}

	data := domain.Permission{
		Permission: entities.Permission{
			Name:        form.Name,
			Description: form.Description,
		},
	}
	if err := uc.store.Create(ctx, &data); err != nil {
		return common.ErrCannotCreateEntity(domain.EntityName, err)
	}

	return nil
}
