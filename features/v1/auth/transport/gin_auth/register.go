package gin_auth

import (
	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/auth/dto"
	"github.com/cesc1802/auth-service/features/v1/auth/usecase"
	"github.com/cesc1802/auth-service/features/v1/user/storage"
	"github.com/cesc1802/auth-service/pkg/hash/md5"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register
// @Summary 	Register
// @Description Register
// @Tags 		Auth
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Success 	200
// @Param 		Permission		body		dto.CreateUserRequest 	true "Create User"
// @Failure 	400 			{object} 	common.AppError
// @Failure 	404 			{object} 	common.AppError
// @Router 		/api/v1/auth/register 			[post]
func Register(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form dto.CreateUserRequest

		if err := c.ShouldBind(&form); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetAppGorm()
		md5Hash := md5.NewMD5Hash()

		store := storage.NewMySqlUserStore(db)
		uc := usecase.NewUseCaseCreateUser(store, md5Hash)

		if err := uc.RegisterUser(c.Request.Context(), &form); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
		return
	}
}
