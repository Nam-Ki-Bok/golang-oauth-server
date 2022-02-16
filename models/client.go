package models

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/models"
	"golang.org/x/oauth2/clientcredentials"
	"infradev-practice/Wade/OAuth2.0-server/database/maria"
	"infradev-practice/Wade/OAuth2.0-server/utils"
)

type OauthClients struct {
	ClientID     string
	ClientSecret string
	ClientIP     string
	GrantType    string
	Scope        string

	Config    *clientcredentials.Config
	SaveModel *models.Client
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

func (c *OauthClients) GetScope() []string {
	return c.Config.Scopes
}

func (c *OauthClients) GetConfig() *clientcredentials.Config {
	return c.Config
}

func (c *OauthClients) GetSaveModel() *models.Client {
	return c.SaveModel
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
		c.SetConfig()
		c.SetSaveModel()
		return true
	}
}

// SetConfig set clientcredentials.Config struct
func (c *OauthClients) SetConfig() {
	scopes := utils.SplitScope(c.Scope)

	cfg := new(clientcredentials.Config)
	cfg.ClientID = c.GetClientID()
	cfg.ClientSecret = c.GetClientSecret()
	cfg.TokenURL = "http://localhost:9096/generate/token"
	cfg.Scopes = scopes

	c.Config = cfg
}

// SetSaveModel set model to store in the client store
func (c *OauthClients) SetSaveModel() {
	model := new(models.Client)
	model.ID = c.GetClientID()
	model.Secret = c.GetClientSecret()
	model.Domain = c.GetClientIP()

	c.SaveModel = model
}
