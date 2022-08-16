package domain

import (
	"fmt"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/entities"
	"github.com/cesc1802/auth-service/features/v1/auth/dto"
)

const (
	EntityName = "user"
)

var (
	ErrUserExisting = common.NewCustomError(nil, "user existing", "ERR_USER_EXISTING")
)

type User struct {
	entities.User
}

func FromUserDto(u *dto.CreateUserRequest) User {
	return User{
		entities.User{
			LoginID:   u.LoginID,
			Password:  u.Password,
			LastName:  u.LastName,
			FirstName: u.FirstName,
			FullName:  fmt.Sprintf("%s %s", u.FirstName, u.LastName),
		},
	}
}
