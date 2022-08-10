package domain

import (
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/entities"
)

const (
	EntityName = "user_role"
)

var (
	ErrNumOfRoleNotEnough = common.NewCustomError(nil, "number of role id not enough",
		"ERR_NUM_OF_ROLE_ID_NOT_ENOUGH")
)

type UserRole struct {
	entities.UserRole
}
