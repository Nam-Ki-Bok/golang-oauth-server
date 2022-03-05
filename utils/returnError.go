package utils

import (
	"github.com/gin-gonic/gin"
)

// ReturnError Return various errors to JSON
func ReturnError(c *gin.Context, code int, err error) {
	c.Error(err)

	recovered := gin.H{
		"code": code,
		"err":  err.Error(),
	}

	panic(recovered)
}
