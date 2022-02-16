package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"

	"golang.org/x/net/context"
	"infradev-practice/Wade/OAuth2.0-server/database/maria"
	"infradev-practice/Wade/OAuth2.0-server/database/redis"
	"infradev-practice/Wade/OAuth2.0-server/models"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"log"
	"net/http"
)

var (
	manager     = manage.NewDefaultManager()
	srv         = server.NewServer(server.NewConfig(), manager)
	clientStore = store.NewClientStore()
)

func init() {
	initManager()

	maria.Connect()
	redis.Connect()
}

func main() {
	r := gin.Default()

	r.GET("/token", publicApiRequestHandler)
	r.POST("/generate/token", func(c *gin.Context) {
		err := srv.HandleTokenRequest(c.Writer, c.Request)
		if err != nil {
			utils.ReturnError(c, http.StatusInternalServerError, "Failed to generate a Token")
		}
	})

	log.Fatal(r.Run(":9096"))
}

func initManager() {
	manager.SetClientTokenCfg(manage.DefaultClientTokenCfg)
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	manager.MapAccessGenerate(generates.NewAccessGenerate())
	manager.MapClientStorage(clientStore)
}

func publicApiRequestHandler(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			return
		}
	}()

	client := models.NewClient(c)
	if client.IsValid() {
		err := SaveClientStore(client)
		if err != nil {
			utils.ReturnError(c, http.StatusInternalServerError, "Failed to save to client store")
		}
	} else {
		utils.ReturnError(c, http.StatusUnauthorized, "Invalid client")
	}

	config := client.GetConfig()
	token, err := config.Token(context.Background())
	if err != nil {
		utils.ReturnError(c, http.StatusInternalServerError, "Failed to generate a token")
	}

	authInfo := models.NewAuthInfo(client, token)
	authInfo.SaveRedis()

	c.JSON(200, authInfo)
}

func SaveClientStore(c *models.OauthClients) error {
	err := clientStore.Set(c.GetClientID(), c.GetSaveModel())
	if err != nil {
		return err
	} else {
		return nil
	}
}
