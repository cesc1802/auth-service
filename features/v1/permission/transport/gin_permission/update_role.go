package gin_permission

import (
	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/features/v1/permission/dto"
	"github.com/cesc1802/auth-service/features/v1/permission/storage"
	"github.com/cesc1802/auth-service/features/v1/permission/usecase"
	"github.com/cesc1802/auth-service/pkg/httpserver/extention"
	"github.com/gin-gonic/gin"
)

// UpdatePermission
// @Summary 	Update Permission
// @Description Update Permission
// @Tags 		Permissions
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Param 		id			path		int		true		"id"
// @Param 		Permission		body		dto.UpdatePermissionRequest 	true "Update Permission"
// @Success 	200
// @Failure 	400 		{object} 	common.AppError
// @Failure 	404 		{object} 	common.AppError
// @Router 		/api/v1/permissions/{id} 		[put]
func UpdatePermission(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		extCtx := extention.NewContextExtension(c)

		PermissionId := extCtx.GetPathParam("id", 0)
		var form dto.UpdatePermissionRequest
		if err := c.ShouldBind(&form); err != nil {
			panic(err)
		}

		db := appCtx.GetAppGorm()
		store := storage.NewMySqlPermissionStore(db)
		uc := usecase.NewUseCaseUpdatePermission(store)

		if err := uc.UpdatePermission(c.Request.Context(), PermissionId, &form); err != nil {
			panic(err)
		}
		extCtx.SimpleSuccessResponse(true)
		return
	}
}
