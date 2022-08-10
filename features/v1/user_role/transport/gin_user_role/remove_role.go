package gin_user_role

import (
	"net/http"

	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/user_role/dto"
	"github.com/cesc1802/auth-service/features/v1/user_role/storage"
	"github.com/cesc1802/auth-service/features/v1/user_role/usecase"
	"github.com/gin-gonic/gin"
)

// RemoveRolesFromUser
// @Summary 	Remove Roles From User
// @Description Remove Roles From User
// @Tags 		User Roles
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Success 	200
// @Param 		Permission		body		dto.RemoveRolesRequest 	true "Remove Roles From User"
// @Failure 	400 			{object} 	common.AppError
// @Failure 	404 			{object} 	common.AppError
// @Router 		/api/v1/user_roles 			[delete]
func RemoveRolesFromUser(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form dto.RemoveRolesRequest
		db := appCtx.GetAppGorm()

		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		userRoleStore := storage.NewMySqlUserRoleStore(db)
		uc := usecase.NewUseCaseDeleteRole(userRoleStore)

		if err := uc.DeleteRole(c.Request.Context(), &form); err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
