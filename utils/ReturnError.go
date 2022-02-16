package utils

import (
	"github.com/gin-gonic/gin"
)

// ReturnError Return various errors to JSON
func ReturnError(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"err": msg,
	})
}
