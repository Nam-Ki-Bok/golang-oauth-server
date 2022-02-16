package models

import (
	"encoding/json"
	"golang.org/x/oauth2"
	"infradev-practice/Wade/OAuth2.0-server/database/redis"
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
