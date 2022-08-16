package usecase

import (
	"context"
	"fmt"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/auth/dto"
	"github.com/cesc1802/auth-service/features/v1/user/domain"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"github.com/cesc1802/auth-service/pkg/hash"
	"github.com/cesc1802/auth-service/pkg/utils/random"
	"gorm.io/gorm"
)

type CreateUserStore interface {
	generic.IFindOneByConditionStore[domain.User]
	generic.ICreateStore[domain.User]
}

type ucCreateUser struct {
	store  CreateUserStore
	hasher hash.Hasher
}

func NewUseCaseCreateUser(store CreateUserStore, hasher hash.Hasher) *ucCreateUser {
	return &ucCreateUser{
		store:  store,
		hasher: hasher,
	}
}

func (uc *ucCreateUser) RegisterUser(ctx context.Context, form *dto.CreateUserRequest) error {
	user, err := uc.store.FindOneByCondition(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("login_id", form.LoginID)
	})

	if err != nil && err != common.ErrRecordNotFound {
		return common.ErrCannotGetEntity(domain.EntityName, err)
	}

	if user != nil {
		return domain.ErrUserExisting
	}

	salt := random.String(50, random.Alphanumeric)
	createUser := domain.FromUserDto(form)
	createUser.Salt = salt
	createUser.Password = uc.hasher.Hash(fmt.Sprintf("%s%s", form.Password, salt))
	createUser.RefreshTokenID = random.String(50, random.Alphanumeric)

	if err := uc.store.Create(ctx, &createUser); err != nil {
		return common.ErrCannotCreateEntity(domain.EntityName, err)
	}

	return nil
}
