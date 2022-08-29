package usecase

import (
	"context"
	"fmt"

	"github.com/cesc1802/auth-service/common"
	rolePermissionDomain "github.com/cesc1802/auth-service/features/v1/role_permissions/domain"
	"github.com/cesc1802/auth-service/pkg/broker"
	"github.com/cesc1802/auth-service/pkg/cache"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

type FindRolePermissionStore interface {
	generic.IFindAllStore[rolePermissionDomain.RolePermission]
}

type cacheUsecase struct {
	store FindRolePermissionStore
	cache cache.ICache
}

func NewCacheUseCase(
	store FindRolePermissionStore,
	cache cache.ICache,
) *cacheUsecase {
	return &cacheUsecase{
		store: store,
		cache: cache,
	}
}

func (uc *cacheUsecase) Handler(delivery amqp.Delivery) {
	var message broker.Message

	type value struct {
		RoleIDs []uint
		UserID  uint
	}
	var val value
	err := message.Unmarshal(delivery.Body, &val)
	if err != nil {
		panic(err)
	}
	rolePermissions, _, _ := uc.store.FindAll(context.Background(), func(db *gorm.DB) *gorm.DB {
		return db.Where("role_id IN (?)", val.RoleIDs)
	})

	var uniquePerms = make(map[uint]bool)
	var permissions []uint
	for _, rolePermission := range rolePermissions {
		if uniquePerms[rolePermission.PermissionID] {
			continue
		}
		uniquePerms[rolePermission.PermissionID] = true
		permissions = append(permissions, rolePermission.PermissionID)
	}
	uc.cache.Set(fmt.Sprintf(common.UserPermissionCacheKey, val.UserID), permissions, common.DefaultCacheExpiration)
}
