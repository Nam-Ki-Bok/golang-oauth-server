package models

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/models"
	"golang.org/x/oauth2/clientcredentials"
	"infradev-practice/Wade/OAuth2.0-server/database/maria"
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

func NewClient(c *gin.Context) *Clients {
	client := new(Clients)

	id := c.GetString("client_id")
	secret := c.GetString("client_secret")
	maria.DB.Where("client_id = ?", id).
		Where("client_secret = ?", secret).
		Where("client_ip = ?", c.ClientIP()).Find(&client)

	client.SetConfig()
	client.SetSaveModel()

	return client
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
