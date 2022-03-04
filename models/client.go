package models

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/models"
	"golang.org/x/oauth2/clientcredentials"
	"infradev-practice/Wade/OAuth2.0-server/database/maria"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"net/http"
	"os"
)

type OauthClients struct {
	ClientID     string
	ClientSecret string
	ClientIP     string
	GrantType    string
	Scope        string

	Config    *clientcredentials.Config `gorm:"-"`
	SaveModel *models.Client            `gorm:"-"`
}

var (
	Client *OauthClients
)

func InitClient(c *gin.Context) {
	id, secret, ok := c.Request.BasicAuth()
	if !ok {
		utils.ReturnError(c, http.StatusUnauthorized, errors.New("authentication information error"))
	}

	if err := utils.CheckID(id); err != nil {
		utils.ReturnError(c, http.StatusUnauthorized, err)
	}

	if err := utils.CheckSecret(secret); err != nil {
		utils.ReturnError(c, http.StatusUnauthorized, err)
	}

	Client = &OauthClients{
		ClientID:     id,
		ClientSecret: utils.GenerateSHA256(secret),
		ClientIP:     c.ClientIP(),
	}
}

// IsValid check client in ACL
func (c *OauthClients) IsValid() bool {
	result := maria.DB.Where("client_id = ?", c.ClientID).
		Where("client_secret = ?", c.ClientSecret).
		Where("client_ip = ?", c.ClientIP).Find(c)

	if result.RowsAffected == 0 {
		return false
	}
	return true
}

// SetConfig set clientcredentials.Config struct
func (c *OauthClients) SetConfig() {
	c.Config = &clientcredentials.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		TokenURL:     os.Getenv("TOKEN_URL"),
		Scopes:       []string{c.Scope},
	}
}

// SetSaveModel set model to store in the client store
func (c *OauthClients) SetSaveModel() {
	c.SaveModel = &models.Client{
		ID:     c.ClientID,
		Secret: c.ClientSecret,
		Domain: c.ClientIP,
	}
}

func (c *OauthClients) GetClientID() string {
	return c.ClientID
}

func (c *OauthClients) SetClientID(id string) {
	c.ClientID = id
}

func (c *OauthClients) GetClientSecret() string {
	return c.ClientSecret
}

func (c *OauthClients) SetClientSecret(secret string) {
	c.ClientSecret = secret
}

func (c *OauthClients) GetClientIP() string {
	return c.ClientIP
}

func (c *OauthClients) SetClientIP(ip string) {
	c.ClientIP = ip
}

func (c *OauthClients) GetGrantType() string {
	return c.GrantType
}

func (c *OauthClients) GetScope() []string {
	return c.Config.Scopes
}

func (c *OauthClients) GetConfig() *clientcredentials.Config {
	return c.Config
}

func (c *OauthClients) GetSaveModel() *models.Client {
	return c.SaveModel
}
