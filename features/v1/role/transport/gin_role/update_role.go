package gin_role

import (
	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/features/v1/role/dto"
	"github.com/cesc1802/auth-service/features/v1/role/storage"
	"github.com/cesc1802/auth-service/features/v1/role/usecase"
	"github.com/cesc1802/auth-service/pkg/httpserver/extention"
	"github.com/gin-gonic/gin"
)

// UpdateRole
// @Summary 	Update Role
// @Description Update Role
// @Tags 		Roles
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Param 		id			path		int		true		"id"
// @Param 		role		body		dto.UpdateRoleRequest 	true "Update Role"
// @Success 	200
// @Failure 	400 		{object} 	common.AppError
// @Failure 	404 		{object} 	common.AppError
// @Router 		/api/v1/roles/{id} 		[put]
func UpdateRole(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		extCtx := extention.NewContextExtension(c)

		roleId := extCtx.GetPathParam("id", 0)
		var form dto.UpdateRoleRequest
		if err := c.ShouldBind(&form); err != nil {
			panic(err)
		}

		db := appCtx.GetAppGorm()
		store := storage.NewMySqlRoleStore(db)
		uc := usecase.NewUseCaseUpdateRole(store)

		if err := uc.UpdateRole(c.Request.Context(), roleId, &form); err != nil {
			panic(err)
		}
		extCtx.SimpleSuccessResponse(true)
		return
	}
}
