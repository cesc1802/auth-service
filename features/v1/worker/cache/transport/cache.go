package transport

import (
	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/features/v1/role_permissions/storage"
	"github.com/cesc1802/auth-service/features/v1/worker/cache/usecase"
)

func LoginCacheWorker(appCtx app_context.AppContext) {
	db := appCtx.GetAppGorm()
	cache := appCtx.GetAppCache()
	subscriber := appCtx.GetSubscriber()

	store := storage.NewMySqlRolePermissionStore(db)
	cacheUc := usecase.NewCacheUseCase(store, cache)

	subscriber.Subscribe("", cacheUc.Handler)
}
