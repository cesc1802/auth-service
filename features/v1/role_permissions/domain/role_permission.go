package domain

import (
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/entities"
)

const (
	EntityName = "role_permission"
)

var (
	ErrNumOfPermissionNotEnough = common.NewCustomError(nil, "number of permission id not enough",
		"ERR_NUM_OF_PERMISSION_ID_NOT_ENOUGH")
)

type RolePermission struct {
	entities.RolePermission
}
