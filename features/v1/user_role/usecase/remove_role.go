package usecase

import (
	"context"
	"errors"

	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/entities"
	"github.com/cesc1802/auth-service/features/v1/user_role/domain"
	"github.com/cesc1802/auth-service/features/v1/user_role/dto"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"gorm.io/gorm"
)

type DeleteRoleStore interface {
	generic.BatchDeleteStore[domain.UserRole]
	generic.CountStore[domain.UserRole]
}

type ucDeleteRole struct {
	store DeleteRoleStore
}

func NewUseCaseDeleteRole(store DeleteRoleStore) *ucDeleteRole {
	return &ucDeleteRole{
		store: store,
	}
}

func (uc *ucDeleteRole) DeleteRole(ctx context.Context, form *dto.RemoveRolesRequest) error {
	if len(form.Roles) == 0 {
		return common.ErrCannotDeleteEntity(domain.EntityName, errors.New("roles cannot be null"))
	}
	roles := make([]uint, len(form.Roles))
	for index, r := range form.Roles {
		roles[index] = r.ID
	}

	filter := func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ? AND role_id in (?)", form.UserID, roles)
	}

	total, err := uc.store.Count(ctx, filter)
	if err != nil {
		return common.ErrCannotGetEntity(domain.EntityName, err)
	}

	if *total < int64(len(roles)) {
		return domain.ErrRolesInvalid
	}

	userRoles := make([]domain.UserRole, len(form.Roles))
	for i, role := range form.Roles {
		userRoles[i] = domain.UserRole{
			UserRole: entities.UserRole{
				UserID: form.UserID,
				RoleID: role.ID,
			},
		}
	}
	if err := uc.store.DeleteByCondition(ctx, userRoles, filter); err != nil {
		return common.ErrCannotDeleteEntity(domain.EntityName, err)
	}

	return nil
}
