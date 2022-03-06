package token

import (
	"context"
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/models"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"net/http"
)

func Request(c *gin.Context) {
	client := models.NewClient(c)

	if err := Cs.Set(client.ClientID, client.SaveModel); err != nil {
		utils.ReturnError(http.StatusBadRequest, err)
	}

	token, err := client.Config.Token(context.Background())
	if err != nil {
		utils.ReturnError(http.StatusUnauthorized, err)
	}

	authCache := models.NewAuthInfo(client, token)
	authCache.SaveRedis()

	c.SecureJSON(http.StatusOK, authCache)
}
