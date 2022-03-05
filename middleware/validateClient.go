package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/models"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"net/http"
)

func ValidateClient(c *gin.Context) {
	models.InitClient(c)
	fmt.Println("여기 변경이 메인엔 없어야지????")
	if ok := models.Client.IsValid(); !ok {
		utils.ReturnError(c, http.StatusUnauthorized, errors.New("invalid client"))
	}

	c.Next()
}
