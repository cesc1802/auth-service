package usecase

import (
	"context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/entities"
	"github.com/cesc1802/auth-service/features/v1/role/domain"
	"github.com/cesc1802/auth-service/features/v1/role/dto"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"gorm.io/gorm"
)

type CreateRoleStore interface {
	generic.IFindOneByConditionStore[domain.Role]
	generic.ICreateStore[domain.Role]
}

type ucCreateRole struct {
	store CreateRoleStore
}

func NewUseCaseCreateRole(store CreateRoleStore) *ucCreateRole {
	return &ucCreateRole{
		store: store,
	}
}

func (uc *ucCreateRole) CreateRole(ctx context.Context, form *dto.CreateRoleRequest) error {
	role, err := uc.store.FindOneByCondition(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("name = ?", form.Name)
	})

	if err != nil && err != common.ErrRecordNotFound(err) {
		return err
	}

	if role != nil {
		return domain.ErrRoleNameIsExisting
	}

	data := domain.Role{
		Role: entities.Role{
			Name:        form.Name,
			Description: form.Description,
		},
	}
	if err := uc.store.Create(ctx, &data); err != nil {
		return err
	}

	return nil
}
