package gin_role

import (
	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/role/dto"
	"github.com/cesc1802/auth-service/features/v1/role/storage"
	"github.com/cesc1802/auth-service/features/v1/role/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateRole
// @Summary 	Create Role
// @Description Create Role
// @Tags 		Roles
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Success 	200
// @Failure 	400 		{object} 	common.AppError
// @Failure 	404 		{object} 	common.AppError
// @Router 		/api/v1/roles 			[post]
func CreateRole(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form dto.CreateRoleRequest

		if err := c.ShouldBind(&form); err != nil {
			panic(err)
		}

		db := appCtx.GetAppGorm()

		store := storage.NewMySqlRoleStore(db)
		uc := usecase.NewUseCaseCreateRole(store)

		if err := uc.CreateRole(c.Request.Context(), &form); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
		return
	}
}
