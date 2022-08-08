package transport

import (
	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/role_permissions/dto"
	"github.com/cesc1802/auth-service/features/v1/role_permissions/storage"
	"github.com/cesc1802/auth-service/features/v1/role_permissions/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateBatchRolePermission(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form dto.CreateRolePermissionRequest
		db := appCtx.GetAppGorm()

		if err := c.ShouldBind(&form); err != nil {

		}

		store := storage.NewMySqlRolePermissionStore(db)
		uc := usecase.NewUseCaseRolePermission(store)

		if err := uc.CreateRolePermission(c.Request.Context(), &form); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
		return
	}
}
