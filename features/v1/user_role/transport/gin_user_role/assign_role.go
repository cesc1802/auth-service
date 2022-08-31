package gin_user_role

import (
	"fmt"
	"net/http"

	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/common"
	storage2 "github.com/cesc1802/auth-service/features/v1/role/storage"
	"github.com/cesc1802/auth-service/features/v1/user_role/dto"
	"github.com/cesc1802/auth-service/features/v1/user_role/storage"
	"github.com/cesc1802/auth-service/features/v1/user_role/usecase"
	"github.com/gin-gonic/gin"
)

// AssignRolesToUser
// @Summary 	Assign Roles To User
// @Description Assign Roles To User
// @Tags 		User Roles
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Success 	200
// @Param 		Permission		body		dto.AssignRolesToUserRequest 	true "Assign Roles to User"
// @Failure 	400 			{object} 	common.AppError
// @Failure 	404 			{object} 	common.AppError
// @Router 		/api/v1/user_roles 			[post]
func AssignRolesToUser(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form dto.AssignRolesToUserRequest
		db := appCtx.GetAppGorm()
		cache := appCtx.GetAppCache()

		if err := c.ShouldBind(&form); err != nil {
			panic(err)
		}

		userRoleStore := storage.NewMySqlUserRoleStore(db)
		roleStore := storage2.NewMySqlRoleStore(db)
		uc := usecase.NewUseCaseAssignRole(roleStore, userRoleStore)

		if err := uc.AssignRoleToUser(c.Request.Context(), &form); err != nil {
			panic(err)
		}

		userRoleCacheKey := fmt.Sprintf(common.UserRoleCacheKey, form.UserID)
		userPermissionCacheKey := fmt.Sprintf(common.UserPermissionCacheKey, form.UserID)
		cache.Delete(userRoleCacheKey)
		cache.Delete(userPermissionCacheKey)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
