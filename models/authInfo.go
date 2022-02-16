package models

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"infradev-practice/Wade/OAuth2.0-server/database/redis"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"time"
)

type AuthInfo struct {
	ClientID    string    `json:"client_id"`
	AccessToken string    `json:"access_token"`
	Scope       []string  `json:"scope"`
	ExpiresIn   time.Time `json:"expires_in"`
}

func NewAuthInfo(c *OauthClients, t *oauth2.Token) *AuthInfo {
	return &AuthInfo{
		ClientID:    c.ClientID,
		AccessToken: t.AccessToken,
		Scope:       c.Config.Scopes,
		ExpiresIn:   t.Expiry,
	}
}

func (a *AuthInfo) SaveRedis() {
	data, _ := json.Marshal(a)
	redis.DB.Set(a.ClientID, data, a.ExpiresIn.Sub(time.Now()))
}

func (a *AuthInfo) IsExists(c *gin.Context) bool {
	id, _ := utils.BindClientInfo(c)

	err := redis.DB.Get(id).Err()
	if err == redis.NilErr {
		return false
	} else {
		data := redis.DB.Get(id)
		_ = json.Unmarshal([]byte(data.Val()), &a)
		return true
	}
}
