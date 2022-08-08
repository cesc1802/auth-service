package domain

import (
	"errors"

	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/entities"
	"github.com/cesc1802/auth-service/errorcode"
)

const (
	EntityName = "user_role"
)

var (
	ErrUserRoleIsAssigned = common.NewCustomError(errors.New("role is assigned to user"), "role is assigned to user", errorcode.ERR_USER_ROLE_IS_ASSIGNED)
)

type UserRole struct {
	entities.UserRole
}
