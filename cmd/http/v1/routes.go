package v1

import (
	"github.com/cesc1802/auth-service/app_context"
	"github.com/cesc1802/auth-service/features/v1/role/transport/gin_role"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
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
		e.POST("/roles", gin_role.CreateRole(appCtx))
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
