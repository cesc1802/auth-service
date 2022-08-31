package usecase

import (
	"context"

	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/role/domain"
	"github.com/cesc1802/auth-service/pkg/broker"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"gorm.io/gorm"
)

type DeleteRoleStore interface {
	generic.IFindOneByConditionStore[domain.Role]
	generic.IDeleteStore[domain.Role]
}

type ucDeleteRole struct {
	store     DeleteRoleStore
	publisher broker.Publisher
}

func NewUseCaseDeleteStore(store DeleteRoleStore, publisher broker.Publisher) *ucDeleteRole {
	return &ucDeleteRole{
		store:     store,
		publisher: publisher,
	}
}

func (uc *ucDeleteRole) DeleteRole(ctx context.Context, id uint) error {
	_, err := uc.store.FindOneByCondition(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	})

	if err != nil && err != common.ErrRecordNotFound {
		return common.ErrCannotGetEntity(domain.EntityName, err)
	}

	if err := uc.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(domain.EntityName, err)
	}

	uc.publisher.Produce(ctx, broker.Message{
		Value: broker.MessageValue{
			RoleIDs: []uint{id},
		},
		Topic: common.DeleteRoleTopic,
	})
	return nil
}
