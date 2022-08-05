package extention

import (
	"github.com/gin-gonic/gin"
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
