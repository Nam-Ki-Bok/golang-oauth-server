package utils

import (
	"infradev-practice/Wade/OAuth2.0-server/database/maria"
	"infradev-practice/Wade/OAuth2.0-server/models"
)

func IsValidClient(c *models.OauthClients) bool {
	err := maria.DB.Where("client_id = ?", c.ClientID).
		Where("client_secret = ?", c.ClientSecret).
		Where("client_ip = ?", c.ClientIP).
		Find(c).Error

	if err != nil {
		return false
	} else {
		return true
	}
}
