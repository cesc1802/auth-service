package usecase

import (
	"context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/permission/domain"
	"github.com/cesc1802/auth-service/features/v1/permission/dto"
	"github.com/cesc1802/auth-service/pkg/database"
	"github.com/cesc1802/auth-service/pkg/database/generic"
	"github.com/cesc1802/auth-service/pkg/paging"
	"github.com/jinzhu/copier"
)

type ListStore interface {
	generic.IFindAllStore[domain.Permission]
}

type ucListPermission struct {
	store ListStore
}

func NewUseCaseListStore(store ListStore) *ucListPermission {
	return &ucListPermission{
		store: store,
	}
}

func (uc *ucListPermission) ListPermission(ctx context.Context, page *paging.Paging,
	filter *dto.Filter) (dto.ListPermissionResponse, error) {
	Permissions, total, err := uc.store.FindAll(ctx, database.QueryPage(page.Offset, page.Limit),
		database.Condition[string]("name", filter.Name),
		database.Condition[string]("description", filter.Description))

	if err != nil {
		return nil, common.ErrCannotListEntity(domain.EntityName, err)
	}

	var results dto.ListPermissionResponse
	if err := copier.Copy(&results, Permissions); err != nil {
		return nil, common.ErrCopyData
	}
	page.Total = *total
	return results, nil

}
