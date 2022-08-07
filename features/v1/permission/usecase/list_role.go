package usecase

import (
	"context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/role/domain"
	"github.com/cesc1802/auth-service/features/v1/role/dto"
	"github.com/cesc1802/auth-service/pkg/database"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"github.com/cesc1802/auth-service/pkg/paging"
	"github.com/jinzhu/copier"
)

type ListStore interface {
	generic.IFindAllStore[domain.Role]
}

type ucListRole struct {
	store ListStore
}

func NewUseCaseListStore(store ListStore) *ucListRole {
	return &ucListRole{
		store: store,
	}
}

func (uc *ucListRole) ListRole(ctx context.Context, page *paging.Paging,
	filter *dto.Filter) (dto.ListRoleResponse, error) {
	roles, total, err := uc.store.FindAll(ctx, database.QueryPage(page.Offset, page.Limit),
		database.Condition[string]("name", filter.Name),
		database.Condition[string]("description", filter.Description))

	if err != nil {
		return nil, common.ErrCannotListEntity(domain.EntityName, err)
	}

	var results dto.ListRoleResponse
	if err := copier.Copy(&results, roles); err != nil {
		return nil, common.ErrCopyData
	}
	page.Total = *total
	return results, nil

}
