package domain

import (
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/entities"
	"github.com/cesc1802/auth-service/errorcode"
	"github.com/pkg/errors"
)

const (
	EntityName = "role"
)

var (
	ErrRoleNameIsExisting = common.NewCustomError(errors.New("role name is existing"), "role name existing", errorcode.ERR_ROLE_EXISTING)
)

type Role struct {
	entities.Role
}
