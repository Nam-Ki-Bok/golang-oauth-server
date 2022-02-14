package utils

import (
	"github.com/gin-gonic/gin"
)

func ReturnError(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"err": msg,
	})
	panic(msg)
}
