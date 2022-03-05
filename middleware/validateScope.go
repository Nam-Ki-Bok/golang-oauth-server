package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/database/maria"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"net/http"
	"strings"
)

type resources struct {
	Resource string
	Scope    int
}

func ValidateScope(c *gin.Context) {
	api := new(resources)
	resource := strings.Split(c.Request.URL.Path, "/")[1]

	err := maria.DB.Where("resource = ?", resource).Find(&api).Error
	if err != nil {
		utils.ReturnError(http.StatusBadRequest, err)
	}

	clientScope := c.GetInt("scope")

	if clientScope < api.Scope {
		utils.ReturnError(http.StatusUnauthorized, errors.New("invalid scope"))
	}

	c.Next()
}
