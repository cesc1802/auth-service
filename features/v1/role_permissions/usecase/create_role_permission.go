package usecase

import (
	"context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/entities"
	permissionDomain "github.com/cesc1802/auth-service/features/v1/Permission/domain"
	"github.com/cesc1802/auth-service/features/v1/role_permissions/domain"
	"github.com/cesc1802/auth-service/features/v1/role_permissions/dto"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type FindAllRoleStore interface {
	generic.IFindAllStore[permissionDomain.Permission]
	generic.CountStore[permissionDomain.Permission]
}
type CreateRolePermissionStore interface {
	generic.IFindAllStore[domain.RolePermission]
	generic.BatchCreateStore[domain.RolePermission]
}

type ucCreateRolePermission struct {
	store           CreateRolePermissionStore
	permissionStore FindAllRoleStore
}

func NewUseCaseRolePermission(store CreateRolePermissionStore, permissionStore FindAllRoleStore) *ucCreateRolePermission {
	return &ucCreateRolePermission{
		store:           store,
		permissionStore: permissionStore,
	}
}

func (uc *ucCreateRolePermission) CreateRolePermission(ctx context.Context, form *dto.CreateRolePermissionRequest) error {
	if len(form.Permissions) <= 0 {
		return common.ErrCannotCreateEntity(domain.EntityName, errors.New("permissions cannot be null"))
	}

	permissionIds := make([]uint, len(form.Permissions))
	for index, p := range form.Permissions {
		permissionIds[index] = p.ID
	}

	total, err := uc.permissionStore.Count(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (?)", permissionIds)
	})

	if err != nil {
		return common.ErrCannotGetEntity(permissionDomain.EntityName, err)
	}

	if *total != int64(len(permissionIds)) {
		return domain.ErrNumOfPermissionNotEnough
	}

	createRolePermission := make([]domain.RolePermission, len(form.Permissions))

	for i, permission := range form.Permissions {
		rolePermission := domain.RolePermission{
			RolePermission: entities.RolePermission{
				RoleID:       form.RoleID,
				PermissionID: permission.ID,
			},
		}
		createRolePermission[i] = rolePermission
	}
	if err := uc.store.BatchCreate(ctx, createRolePermission); err != nil {
		return common.ErrCannotCreateEntity("", err)
	}

	return nil
}
