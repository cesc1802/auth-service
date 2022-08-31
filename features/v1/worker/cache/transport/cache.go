package transport

import (
	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/features/v1/role_permissions/storage"
	storage2 "github.com/cesc1802/auth-service/features/v1/user_role/storage"
	"github.com/cesc1802/auth-service/features/v1/worker/cache/usecase"
)

func LoginCacheWorker(appCtx app_context.AppContext) {
	db := appCtx.GetAppGorm()
	cache := appCtx.GetAppCache()
	subscriber := appCtx.GetSubscriber()

	store := storage.NewMySqlRolePermissionStore(db)
	store2 := storage2.NewMySqlUserRoleStore(db)
	cacheUc := usecase.NewCacheUseCase(store, store2, cache)

	subscriber.Subscribe(cacheUc.Handler)
}
