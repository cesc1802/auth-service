package usecase

import (
	"context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/role/domain"
	"github.com/cesc1802/auth-service/features/v1/role/dto"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UpdateStore interface {
	generic.IFindOneByConditionStore[domain.Role]
	generic.IUpdateStore[domain.Role]
}

type ucUpdateRole struct {
	store UpdateStore
}

func NewUseCaseUpdateRole(store UpdateStore) *ucUpdateRole {
	return &ucUpdateRole{
		store: store,
	}
}

func (uc *ucUpdateRole) UpdateRole(ctx context.Context, id uint, form *dto.UpdateRoleRequest) error {
	role, err := uc.store.FindOneByCondition(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})

	if err != nil {
		return common.ErrCannotGetEntity(domain.EntityName, err)
	}

	if err := copier.Copy(role, form); err != nil {
		return common.ErrCopyData
	}
	if err := uc.store.Update(ctx, role); err != nil {
		return common.ErrCannotUpdateEntity(domain.EntityName, err)
	}

	return nil
}
