package tokenprovider

import (
	"github.com/cesc1802/auth-service/common"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ERR_TOKEN_NOT_FOUND",
	)
	ErrEncodingToken = common.NewCustomError(errors.New("error encoding the token"),
		"error encoding the token",
		"ERR_ENCODING_TOKEN",
	)
	ErrInvalidToken = common.NewCustomError(errors.New("invalid token provided"),
		"invalid token provided",
		"ERR_INVALID_TOKEN",
	)
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  uint      `json:"expiry"`
}

type TokenPayload struct {
	UserId         uint   `json:"user_id"`
	RefreshTokenId string `json:"refresh_token_id,omitempty"`
	Roles          []uint `json:"roles"`
}
