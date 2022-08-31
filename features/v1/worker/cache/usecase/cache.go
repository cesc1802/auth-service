package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cesc1802/auth-service/common"
	rolePermissionDomain "github.com/cesc1802/auth-service/features/v1/role_permissions/domain"
	userRoleDomain "github.com/cesc1802/auth-service/features/v1/user_role/domain"
	"github.com/cesc1802/auth-service/pkg/broker"
	"github.com/cesc1802/auth-service/pkg/cache"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

type FindRolePermissionStore interface {
	generic.IFindAllStore[rolePermissionDomain.RolePermission]
}

type FindUserRoleStore interface {
	generic.IFindAllStore[userRoleDomain.UserRole]
}

type cacheUsecase struct {
	rolePermstore FindRolePermissionStore
	userRoleStore FindUserRoleStore
	cache         cache.ICache
}

func NewCacheUseCase(
	rolePermstore FindRolePermissionStore,
	userRoleStore FindUserRoleStore,
	cache cache.ICache,
) *cacheUsecase {
	return &cacheUsecase{
		rolePermstore: rolePermstore,
		userRoleStore: userRoleStore,
		cache:         cache,
	}
}

func (uc *cacheUsecase) Handler(delivery amqp.Delivery) {
	var message broker.Message
	err := json.Unmarshal(delivery.Body, &message)
	if err != nil {
		panic(err)
	}
	switch message.Topic {
	case common.LoginTopic:
		uc.loginHandler(message.Value)
	case common.AssignRolePermissionTopic, common.RemoveRolePermissionTopic, common.DeleteRoleTopic:
		uc.roleHandler(message.Value)
	case common.DeletePermissionTopic:
		uc.deletePermissionHandler(message.Value)
	}
}

func (uc *cacheUsecase) loginHandler(val broker.MessageValue) {
	rolePermissions, _, _ := uc.rolePermstore.FindAll(context.Background(), func(db *gorm.DB) *gorm.DB {
		return db.Where("role_id IN (?)", val.RoleIDs)
	})

	if len(rolePermissions) == 0 {
		return
	}

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

func (uc *cacheUsecase) roleHandler(val broker.MessageValue) {
	userRoles, _, _ := uc.userRoleStore.FindAll(context.Background(), func(db *gorm.DB) *gorm.DB {
		return db.Where("role_id IN (?)", val.RoleIDs)
	})

	if len(userRoles) == 0 {
		return
	}

	var cacheKeys []string
	for _, userRole := range userRoles {
		cacheKeys = append(cacheKeys, fmt.Sprintf(common.UserPermissionCacheKey, userRole.UserID))
	}

	uc.cache.Delete(cacheKeys...)
}

func (uc *cacheUsecase) deletePermissionHandler(val broker.MessageValue) {
	rolePerms, _, _ := uc.rolePermstore.FindAll(context.Background(), func(db *gorm.DB) *gorm.DB {
		return db.Where("permission_id IN (?)", val.PermissionIDs)
	})

	if len(rolePerms) == 0 {
		return
	}

	var roleIDs []uint
	for _, rolePerm := range rolePerms {
		roleIDs = append(roleIDs, rolePerm.RoleID)
	}

	userRoles, _, _ := uc.userRoleStore.FindAll(context.Background(), func(db *gorm.DB) *gorm.DB {
		return db.Where("role_id IN (?)", roleIDs)
	})

	if len(userRoles) == 0 {
		return
	}

	var cacheKeys []string
	for _, userRole := range userRoles {
		cacheKeys = append(cacheKeys, fmt.Sprintf(common.UserPermissionCacheKey, userRole.UserID))
	}

	uc.cache.Delete(cacheKeys...)
}
