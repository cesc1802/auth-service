package gin_permission

import (
	"net/http"

	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/permission/dto"
	"github.com/cesc1802/auth-service/features/v1/permission/storage"
	"github.com/cesc1802/auth-service/features/v1/permission/usecase"
	"github.com/cesc1802/auth-service/pkg/paging"
	"github.com/gin-gonic/gin"
)

// ListPermission
// @Summary 	List Permission
// @Description List Permission
// @Tags 		Permissions
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Param 		offset			query		int		true		"offset"
// @Param 		limit			query		int		false		"limit"
// @Param 		name			query		string 	false		"Permission name"
// @Param 		description		query		string 	false		"Permission description"
// @Success 	200
// @Failure 	400 		{object} 	common.AppError
// @Failure 	404 		{object} 	common.AppError
// @Router 		/api/v1/permissions 		[get]
func ListPermission(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var page paging.Paging
		var filter dto.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.ShouldBind(&page); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		page.Fulfill()

		db := appCtx.GetAppGorm()
		store := storage.NewMySqlPermissionStore(db)
		uc := usecase.NewUseCaseListStore(store)

		Permissions, err := uc.ListPermission(c.Request.Context(), &page, &filter)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(Permissions, page, filter))
		return
	}
}
