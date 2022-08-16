package entities

import "github.com/cesc1802/auth-service/common"

type User struct {
	common.BaseModel
	FullName       string `gorm:"column:full_name"`
	LastName       string `gorm:"column:last_name"`
	FirstName      string `gorm:"column:first_name"`
	LoginID        string `gorm:"column:login_id"`
	Password       string `gorm:"column:password"`
	Salt           string `gorm:"column:salt"`
	RefreshTokenID string `gorm:"column:refresh_token_id"`
}

func (r User) TableName() string {
	return "users"
}

func (r User) EntityName() string {
	return "user"
}
