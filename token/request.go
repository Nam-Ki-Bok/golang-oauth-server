package token

import (
	"context"
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/models"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"net/http"
)

func Request(c *gin.Context) {
	models.Client.SetSaveModel()
	err := Cs.Set(models.Client.GetClientID(), models.Client.GetSaveModel())
	if err != nil {
		utils.ReturnError(http.StatusBadRequest, err)
	}

	models.Client.SetConfig()
	cfg := models.Client.GetConfig()
	token, err := cfg.Token(context.Background())
	if err != nil {
		utils.ReturnError(http.StatusUnauthorized, err)
	}

	authCache := models.NewAuthInfo(models.Client, token)
	authCache.SaveRedis()

	c.SecureJSON(http.StatusOK, authCache)
}
