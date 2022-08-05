package gin_role

import (
	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/common"
	"github.com/cesc1802/auth-service/features/v1/role/dto"
	"github.com/cesc1802/auth-service/features/v1/role/storage"
	"github.com/cesc1802/auth-service/features/v1/role/usecase"
	"github.com/cesc1802/auth-service/pkg/paging"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListRole
// @Summary 	List Role
// @Description List Role
// @Tags 		Roles
// @Accept  	json
// @Produce  	json
// @Security 	ApiKeyAuth
// @Param 		offset			query		int		true		"offset"
// @Param 		limit			query		int		false		"limit"
// @Param 		name			query		string 	false		"role name"
// @Param 		description		query		string 	false		"role description"
// @Success 	200
// @Failure 	400 		{object} 	common.AppError
// @Failure 	404 		{object} 	common.AppError
// @Router 		/api/v1/roles 		[get]
func ListRole(appCtx app_context.AppContext) gin.HandlerFunc {
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
		store := storage.NewMySqlRoleStore(db)
		uc := usecase.NewUseCaseListStore(store)

		roles, err := uc.ListRole(c.Request.Context(), &page, &filter)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(roles, page, filter))
		return
	}
}
