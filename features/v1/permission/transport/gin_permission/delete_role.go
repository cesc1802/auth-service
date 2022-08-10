package gin_permission

import (
	"net/http"

	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/permission/storage"
	"github.com/cesc1802/auth-service/features/v1/permission/usecase"
	"github.com/cesc1802/auth-service/pkg/httpserver/extention"
	"github.com/gin-gonic/gin"
)

// DeletePermission
// @Summary 	Delete Permission
// @Description Delete Permission
// @Tags 		Permissions
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Param 		id			path		int		true		"id"
// @Success 	200
// @Failure 	400 		{object} 	common.AppError
// @Failure 	404 		{object} 	common.AppError
// @Router 		/api/v1/permissions/{id} 		[delete]
func DeletePermission(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetAppGorm()
		extension := extention.NewContextExtension(c)
		PermissionId := extension.GetPathParam("id", 0)

		store := storage.NewMySqlPermissionStore(db)
		uc := usecase.NewUseCaseDeleteStore(store)

		if err := uc.DeletePermission(c.Request.Context(), PermissionId); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
		return
	}
}
