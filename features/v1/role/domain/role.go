package domain

import (
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/entities"
	"github.com/pkg/errors"
)

const (
	EntityName = "role"
)

var (
	ErrRoleNameIsExisting = common.NewCustomError(errors.New("role name is existing"), "", "")
)

type Role struct {
	entities.Role
}
