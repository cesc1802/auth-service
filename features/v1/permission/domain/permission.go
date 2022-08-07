package domain

import (
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/entities"
	"github.com/cesc1802/auth-service/errorcode"
	"github.com/pkg/errors"
)

const (
	EntityName = "Permission"
)

var (
	ErrPermissionNameIsExisting = common.NewCustomError(errors.New("permission name is existing"),
		"permission name existing", errorcode.ERR_PERMISSION_EXISTING)
)

type Permission struct {
	entities.Permission
}
