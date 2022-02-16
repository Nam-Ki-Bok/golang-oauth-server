package token

import (
	"context"
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/models"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"log"
	"net/http"
)

func Request(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	client := models.NewClient(c)
	if client.IsValid() {
		if authInfo.IsExists(client) {
			c.JSON(http.StatusOK, authInfo)
			return
		}

		err := saveClientStore(client)
		if err != nil {
			utils.ReturnError(c, http.StatusInternalServerError, "Failed to save to client store")
		}
	} else {
		utils.ReturnError(c, http.StatusUnauthorized, "Invalid client")
	}

	cfg := client.GetConfig()
	token, err := cfg.Token(context.Background())
	if err != nil {
		utils.ReturnError(c, http.StatusInternalServerError, err.Error())
	}

	authInfo = models.NewAuthInfo(client, token)
	authInfo.SaveRedis()

	c.JSON(http.StatusOK, authInfo)
	return
}

func saveClientStore(c *models.OauthClients) error {
	err := cs.Set(c.GetClientID(), c.GetSaveModel())
	if err != nil {
		return err
	} else {
		return nil
	}
}
