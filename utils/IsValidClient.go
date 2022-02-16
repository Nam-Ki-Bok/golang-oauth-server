package utils

import (
	"infradev-practice/Wade/OAuth2.0-server/database/maria"
	"infradev-practice/Wade/OAuth2.0-server/models"
)

func IsValidClient(id, pw string) bool {
	err := maria.DB.Where("client_id = ?", id).
		Where("client_secret = ?", pw).
		Where("client_ip = ?", "1.1.1.1"). // 1.1.1.1 -> c.ClientIP()
		Find(&models.OauthClients{}).Error

	if err != nil {
		return false
	} else {
		return true
	}
}
