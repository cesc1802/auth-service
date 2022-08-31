package gin_auth

import (
	"net/http"

	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/auth/dto"
	"github.com/cesc1802/auth-service/features/v1/auth/usecase"
	storage3 "github.com/cesc1802/auth-service/features/v1/role_permissions/storage"
	"github.com/cesc1802/auth-service/features/v1/user/storage"
	storage2 "github.com/cesc1802/auth-service/features/v1/user_role/storage"
	"github.com/cesc1802/auth-service/pkg/hash/md5"
	"github.com/gin-gonic/gin"
)

// Login
// @Summary 	Login
// @Description Login
// @Tags 		Auth
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Success 	200
// @Param 		Permission		body		dto.LoginUserRequest 	true "Create User"
// @Failure 	400 			{object} 	common.AppError
// @Failure 	404 			{object} 	common.AppError
// @Router 		/api/v1/auth/login 			[post]
func Login(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form dto.LoginUserRequest

		if err := c.ShouldBind(&form); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetAppGorm()

		store := storage.NewMySqlUserStore(db)
		userRoleStore := storage2.NewMySqlUserRoleStore(db)
		rolePermissionStore := storage3.NewMySqlRolePermissionStore(db)
		hasher := md5.NewMD5Hash()
		atProvider := appCtx.GetATProvider()
		rtProvider := appCtx.GetRTProvider()
		publisher := appCtx.GetPublisher()

		uc := usecase.NewUseCaseLogin(store, userRoleStore, rolePermissionStore, hasher, atProvider, rtProvider, publisher)

		data, err := uc.Login(c.Request.Context(), &form)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
		return
	}
}
