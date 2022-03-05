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

type Clients struct {
	ClientID     string
	ClientSecret string
	ClientIP     string
	GrantType    string
	Scope        string

	Config    *clientcredentials.Config `gorm:"-"`
	SaveModel *models.Client            `gorm:"-"`
}

var (
	Client *Clients
)

func InitClient(c *gin.Context) {
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

	Client = &Clients{
		ClientID:     id,
		ClientSecret: utils.GenerateSHA256(secret),
		ClientIP:     c.ClientIP(),
	}
}

// IsValid check client in ACL
func (c *Clients) IsValid() bool {
	result := maria.DB.Where("client_id = ?", c.ClientID).
		Where("client_secret = ?", c.ClientSecret).
		Where("client_ip = ?", c.ClientIP).Find(c)

	if result.RowsAffected == 0 {
		return false
	}
	return true
}

// SetConfig set clientcredentials.Config struct
func (c *Clients) SetConfig() {
	c.Config = &clientcredentials.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		TokenURL:     os.Getenv("TOKEN_URL"),
		Scopes:       []string{c.Scope},
	}
}

// SetSaveModel set model to store in the client store
func (c *Clients) SetSaveModel() {
	c.SaveModel = &models.Client{
		ID:     c.ClientID,
		Secret: c.ClientSecret,
		Domain: c.ClientIP,
	}
}

func (c *Clients) GetClientID() string {
	return c.ClientID
}

func (c *Clients) SetClientID(id string) {
	c.ClientID = id
}

func (c *Clients) GetClientSecret() string {
	return c.ClientSecret
}

func (c *Clients) SetClientSecret(secret string) {
	c.ClientSecret = secret
}

func (c *Clients) GetClientIP() string {
	return c.ClientIP
}

func (c *Clients) SetClientIP(ip string) {
	c.ClientIP = ip
}

func (c *Clients) GetGrantType() string {
	return c.GrantType
}

func (c *Clients) GetScope() []string {
	return c.Config.Scopes
}

func (c *Clients) GetConfig() *clientcredentials.Config {
	return c.Config
}

func (c *Clients) GetSaveModel() *models.Client {
	return c.SaveModel
}
