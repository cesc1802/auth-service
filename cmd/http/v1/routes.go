package v1

import (
	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/docs"
	"github.com/cesc1802/auth-service/features/v1/permission/transport/gin_permission"
	"github.com/cesc1802/auth-service/features/v1/role/transport/gin_role"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func publicRoute(appCtx app_context.AppContext) func(e *gin.RouterGroup) {
	return func(e *gin.RouterGroup) {
		e.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
			return
		})
	}
}

func privateRoute(appCtx app_context.AppContext) func(e *gin.RouterGroup) {
	return func(e *gin.RouterGroup) {
		// Roles
		roles := e.Group("/roles")
		{
			roles.GET("", gin_role.ListRole(appCtx))
			roles.POST("", gin_role.CreateRole(appCtx))
			roles.PUT("/:id", gin_role.UpdateRole(appCtx))
			roles.DELETE("/:id", gin_role.DeleteRole(appCtx))
		}
		permissions := e.Group("/permissions")
		{
			permissions.GET("", gin_permission.ListPermission(appCtx))
			permissions.POST("", gin_permission.CreatePermission(appCtx))
			permissions.PUT("/:id", gin_permission.UpdatePermission(appCtx))
			permissions.DELETE("/:id", gin_permission.DeletePermission(appCtx))
		}
	}
}

func swaggerRoute(appCtx app_context.AppContext) func(e *gin.RouterGroup) {
	return func(e *gin.RouterGroup) {
		docs.SwaggerInfo.BasePath = ""
		e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func SetupRoute(appCtx app_context.AppContext) func(e *gin.Engine) {
	return func(e *gin.Engine) {
		v1 := e.Group("/api/v1")

		publicRoute(appCtx)(v1)
		privateRoute(appCtx)(v1)
		swaggerRoute(appCtx)(v1)
	}
}
