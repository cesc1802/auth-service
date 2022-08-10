package usecase

import (
	"context"
	"errors"

	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/entities"
	"github.com/cesc1802/auth-service/features/v1/role/domain"
	ur_domain "github.com/cesc1802/auth-service/features/v1/user_role/domain"
	"github.com/cesc1802/auth-service/features/v1/user_role/dto"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"gorm.io/gorm"
)

type FindRoleStore interface {
	generic.IFindAllStore[domain.Role]
	generic.CountStore[domain.Role]
}

type AssignRoleStore interface {
	generic.IFindAllStore[ur_domain.UserRole]
	generic.BatchCreateStore[ur_domain.UserRole]
}

type ucAssignRole struct {
	roleStore     FindRoleStore
	userRoleStore AssignRoleStore
}

func NewUseCaseAssignRole(roleStore FindRoleStore, userRoleStore AssignRoleStore) *ucAssignRole {
	return &ucAssignRole{
		roleStore:     roleStore,
		userRoleStore: userRoleStore,
	}
}

func (uc *ucAssignRole) AssignRoleToUser(ctx context.Context, form *dto.AssignRolesToUserRequest) error {
	if len(form.Roles) == 0 {
		return common.ErrCannotCreateEntity(domain.EntityName, errors.New("roles cannot be null"))
	}
	roles := make([]uint, len(form.Roles))
	for index, r := range form.Roles {
		roles[index] = r.ID
	}

	total, err := uc.roleStore.Count(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (?)", roles)
	})
	if err != nil {
		return common.ErrCannotGetEntity(domain.EntityName, err)
	}

	if *total != int64(len(roles)) {
		return ur_domain.ErrNumOfRoleNotEnough
	}

	userRoles := make([]ur_domain.UserRole, len(form.Roles))
	for i, role := range form.Roles {
		userRoles[i] = ur_domain.UserRole{
			UserRole: entities.UserRole{
				UserID: form.UserID,
				RoleID: role.ID,
			},
		}
	}
	if err := uc.userRoleStore.BatchCreate(ctx, userRoles); err != nil {
		return common.ErrCannotCreateEntity("", err)
	}

	return nil
}
