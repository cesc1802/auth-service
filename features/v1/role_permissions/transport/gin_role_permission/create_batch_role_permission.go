package gin_role_permission

import (
	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/common"
	storage2 "github.com/cesc1802/auth-service/features/v1/permission/storage"
	"github.com/cesc1802/auth-service/features/v1/role_permissions/dto"
	"github.com/cesc1802/auth-service/features/v1/role_permissions/storage"
	"github.com/cesc1802/auth-service/features/v1/role_permissions/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateBatchRolePermission
// @Summary 	Create Batch Role Permission
// @Description Create Batch Role Permission
// @Tags 		Role Permissions
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Success 	200
// @Param 		Permission		body		dto.CreateRolePermissionRequest 	true "Create Permission"
// @Failure 	400 			{object} 	common.AppError
// @Failure 	404 			{object} 	common.AppError
// @Router 		/api/v1/role_permissions 			[post]
func CreateBatchRolePermission(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form dto.CreateRolePermissionRequest
		db := appCtx.GetAppGorm()

		if err := c.ShouldBind(&form); err != nil {
			panic(err)
		}

		store := storage.NewMySqlRolePermissionStore(db)
		permissionStore := storage2.NewMySqlPermissionStore(db)
		uc := usecase.NewUseCaseRolePermission(store, permissionStore)

		if err := uc.CreateRolePermission(c.Request.Context(), &form); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
		return
	}
}
