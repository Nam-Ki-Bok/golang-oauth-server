package models

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"infradev-practice/Wade/OAuth2.0-server/database/redis"
	"time"
)

type AuthInfo struct {
	ClientID    string   `json:"client_id"`
	AccessToken string   `json:"access_token"`
	Scope       []string `json:"scope"`
	CreatedIn   string   `json:"created_in"`
	ExpiresIn   string   `json:"expires_in"`
}

var defaultDateFormat = "2006-01-02 15:04:05"
var ctx = context.Background()

func NewAuthInfo(c *Clients, t *oauth2.Token) *AuthInfo {
	return &AuthInfo{
		ClientID:    c.ClientID,
		AccessToken: t.AccessToken,
		Scope:       c.Config.Scopes,
		CreatedIn:   time.Now().Format(defaultDateFormat),
		ExpiresIn:   t.Expiry.Format(defaultDateFormat),
	}
}

func (a *AuthInfo) SaveRedis() {
	data, _ := json.Marshal(a)

	createdIn, _ := time.Parse(defaultDateFormat, a.CreatedIn)
	expiresIn, _ := time.Parse(defaultDateFormat, a.ExpiresIn)

	redis.DB.Set(ctx, a.ClientID, data, expiresIn.Sub(createdIn))
}
