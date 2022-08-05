package gin_role

import (
	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/role/storage"
	"github.com/cesc1802/auth-service/features/v1/role/usecase"
	"github.com/cesc1802/auth-service/pkg/httpserver/extention"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteRole
// @Summary 	Delete Role
// @Description Delete Role
// @Tags 		Roles
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Param 		id			path		int		true		"id"
// @Success 	200
// @Failure 	400 		{object} 	common.AppError
// @Failure 	404 		{object} 	common.AppError
// @Router 		/api/v1/roles/{id} 		[delete]
func DeleteRole(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetAppGorm()
		extension := extention.NewContextExtension(c)
		roleId := extension.GetPathParam("id", 0)

		store := storage.NewMySqlRoleStore(db)
		uc := usecase.NewUseCaseDeleteStore(store)

		if err := uc.DeleteRole(c.Request.Context(), roleId); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
		return
	}
}
