package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"net/http"
)

func Publish(c *gin.Context) {
	c.Next()

	msg, ok := c.Get("msg")
	if !ok {
		utils.ReturnError(http.StatusBadRequest, errors.New("message dose not exist"))
	}

	// execute kafka publish
	fmt.Printf("kafka publish : %s\n", msg)

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
