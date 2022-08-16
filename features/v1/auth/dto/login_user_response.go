package dto

import "github.com/cesc1802/auth-service/pkg/tokenprovider"

type LoginUserResponse struct {
	AccessToken  tokenprovider.Token `json:"access_token"`
	RefreshToken tokenprovider.Token `json:"refresh_token"`
}
