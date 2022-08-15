package gin_user_role

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/user_role/domain"
	"github.com/cesc1802/auth-service/features/v1/user_role/storage"
	"github.com/cesc1802/auth-service/features/v1/user_role/usecase"
	"github.com/cesc1802/auth-service/pkg/httpserver/extention"
	"github.com/gin-gonic/gin"
)

// GetRolesByUserID
// @Summary 	Get Roles By User ID
// @Description Get Roles By User ID
// @Tags 		User Roles
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Param 		user_id		path		int		true		"user_id"
// @Success 	200
// @Failure 	400 		{object} 	common.AppError
// @Failure 	404 		{object} 	common.AppError
// @Router 		/api/v1/user_roles/{user_id} 			[get]
func GetRolesByUserID(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctxExtension := extention.NewContextExtension(c)
		userID := ctxExtension.GetPathParam("user_id", 0)

		cache := appCtx.GetAppCache()
		userRoleCacheKey := fmt.Sprintf(common.UserRoleCacheKey, userID)
		cachedRoles, err := cache.Get(userRoleCacheKey)
		if err != nil {
			if !errors.Is(err, common.ErrNoCacheKeyFound) && !errors.Is(err, common.ErrCacheKeyExpired) {
				panic(err)
			}
		}

		var roles []domain.UserRole
		byteVal, _ := json.Marshal(cachedRoles)
		if err == nil {
			err = json.Unmarshal(byteVal, &roles)
			if err == nil {
				c.JSON(http.StatusOK, common.SimpleSuccessResponse(roles))
				return
			}
		}

		db := appCtx.GetAppGorm()
		store := storage.NewMySqlUserRoleStore(db)
		uc := usecase.NewUseCaseGetRoles(store)

		roles, err = uc.GetRoles(c.Request.Context(), userID)
		if err != nil {
			panic(err)
		}

		cache.Set(userRoleCacheKey, roles, time.Hour*24)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(roles))
	}
}
