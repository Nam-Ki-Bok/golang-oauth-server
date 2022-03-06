package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/database/maria"
	"infradev-practice/Wade/OAuth2.0-server/models"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"net/http"
)

func ValidateClient(c *gin.Context) {
	id, secret, ok := c.Request.BasicAuth()
	if !ok {
		utils.ReturnError(http.StatusUnauthorized, errors.New("authentication information error"))
	}

	if err := utils.CheckID(id); err != nil {
		utils.ReturnError(http.StatusUnauthorized, err)
	}

	if err := utils.CheckSecret(secret); err != nil {
		utils.ReturnError(http.StatusUnauthorized, err)
	}

	secret = utils.GenerateSHA256(secret)
	result := maria.DB.Where("client_id = ?", id).
		Where("client_secret = ?", secret).
		Where("client_ip = ?", c.ClientIP()).Find(&models.Clients{})
	if result.RowsAffected == 0 {
		utils.ReturnError(http.StatusUnauthorized, errors.New("invalid client"))
	}

	c.Set("client_id", id)
	c.Set("client_secret", secret)
	c.Next()
}
