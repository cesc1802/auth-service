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
	ErrUserExisting      = common.NewCustomError(nil, "user existing", "ERR_USER_EXISTING")
	ErrUserBlocked       = common.NewCustomError(nil, "user has been block", "ERR_USER_BLOCKED")
	ErrInvalidCredential = common.NewCustomError(nil, "invalid credentials", "ERR_INVALID_CREDENTIAL")
)

type User struct {
	entities.User
}

func (u User) IsActive() bool {
	return u.Status == 1
}

func (u User) IsBlocked() bool {
	return u.Status == 0
}

func (u User) InvalidPassword(hashPassword string) bool {
	return u.Password != hashPassword
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
