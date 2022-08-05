package extention

import (
	"github.com/cesc1802/auth-service/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type contextExtension struct {
	c *gin.Context
}

func NewContextExtension(c *gin.Context) *contextExtension {
	return &contextExtension{
		c: c,
	}
}

func (extension *contextExtension) GetPathParam(key string, fallback uint) uint {
	value, err := strconv.Atoi(extension.c.Param(key))
	if err != nil {
		return fallback
	}
	return uint(value)
}

func (extension *contextExtension) SimpleSuccessResponse(data interface{}) {
	extension.c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
}
