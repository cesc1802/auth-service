package gin_permission

import (
	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/permission/dto"
	"github.com/cesc1802/auth-service/features/v1/permission/storage"
	"github.com/cesc1802/auth-service/features/v1/permission/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreatePermission
// @Summary 	Create Permission
// @Description Create Permission
// @Tags 		Permissions
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Success 	200
// @Param 		Permission		body		dto.CreatePermissionRequest 	true "Create Permission"
// @Failure 	400 			{object} 	common.AppError
// @Failure 	404 			{object} 	common.AppError
// @Router 		/api/v1/permissions 			[post]
func CreatePermission(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form dto.CreatePermissionRequest

		if err := c.ShouldBind(&form); err != nil {
			panic(err)
		}

		db := appCtx.GetAppGorm()

		store := storage.NewMySqlPermissionStore(db)
		uc := usecase.NewUseCaseCreatePermission(store)

		if err := uc.CreatePermission(c.Request.Context(), &form); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
		return
	}
}
