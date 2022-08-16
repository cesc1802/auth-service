package usecase

import (
	"context"
	"fmt"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/auth/dto"
	"github.com/cesc1802/auth-service/features/v1/user/domain"
	userRoleDomain "github.com/cesc1802/auth-service/features/v1/user_role/domain"
	"github.com/cesc1802/auth-service/pkg/database"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"github.com/cesc1802/auth-service/pkg/hash"
	"github.com/cesc1802/auth-service/pkg/tokenprovider"
	"gorm.io/gorm"
)

type FindUserStore interface {
	generic.IFindOneByConditionStore[domain.User]
}

type FindUserRoleStore interface {
	generic.IFindAllStore[userRoleDomain.UserRole]
}

type ucLoginUser struct {
	store              FindUserStore
	userRoleStore      FindUserRoleStore
	hasher             hash.Hasher
	tokProvider        tokenprovider.Provider
	refreshTokProvider tokenprovider.Provider
}

func NewUseCaseLogin(store FindUserStore, userRoleStore FindUserRoleStore,
	hasher hash.Hasher,
	tokProvider tokenprovider.Provider, refreshTokProvider tokenprovider.Provider) *ucLoginUser {
	return &ucLoginUser{
		store:              store,
		userRoleStore:      userRoleStore,
		hasher:             hasher,
		tokProvider:        tokProvider,
		refreshTokProvider: refreshTokProvider,
	}
}

func (uc *ucLoginUser) Login(ctx context.Context, form *dto.LoginUserRequest) (*dto.LoginUserResponse, error) {
	var result dto.LoginUserResponse

	user, err := uc.store.FindOneByCondition(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("login_id = ?", form.LoginID)
	})

	if err != nil {
		return nil, common.ErrCannotGetEntity(domain.EntityName, err)
	}

	if user.IsBlocked() {
		return nil, domain.ErrUserBlocked
	}

	if user.InvalidPassword(uc.hasher.Hash(fmt.Sprintf("%s%s", form.Password, user.Salt))) {
		return nil, domain.ErrInvalidCredential
	}

	roles, _, err := uc.userRoleStore.FindAll(ctx, func(db *gorm.DB) *gorm.DB {
		return database.ActiveRecord(db).Where("user_id = ?", user.ID)
	})

	if err != nil && err != common.ErrRecordNotFound {
		return nil, common.ErrCannotGetEntity(userRoleDomain.EntityName, err)
	}

	roleIds := make([]uint, len(roles))
	for idx, role := range roles {
		roleIds[idx] = role.RoleID
	}

	accessTok, err := uc.tokProvider.Generate(tokenprovider.TokenPayload{
		UserId: user.ID,
		Roles:  roleIds,
	})

	refreshTok, err := uc.tokProvider.Generate(tokenprovider.TokenPayload{
		UserId:         user.ID,
		RefreshTokenId: user.RefreshTokenID,
		Roles:          roleIds,
	})

	result.AccessToken = *accessTok
	result.RefreshToken = *refreshTok
	return &result, nil
}
