package models

import (
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/database/maria"
	"infradev-practice/Wade/OAuth2.0-server/utils"
)

type OauthClients struct {
	ClientID     string
	ClientSecret string
	ClientIP     string
	GrantType    string
	Scope        string
}

// NewClient create to client model instance
func NewClient(c *gin.Context) *OauthClients {
	id, secret := utils.BindClientInfo(c)

	return &OauthClients{
		ClientID:     id,
		ClientSecret: secret,
		ClientIP:     c.ClientIP(),
	}
}

func (c *OauthClients) GetClientID() string {
	return c.ClientID
}

func (c *OauthClients) GetClientSecret() string {
	return c.ClientSecret
}

func (c *OauthClients) GetClientIP() string {
	return c.ClientIP
}

func (c *OauthClients) GetGrantType() string {
	return c.GrantType
}

func (c *OauthClients) GetScope() string {
	return c.Scope
}

// IsValid check client in ACL
func (c *OauthClients) IsValid() bool {
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
