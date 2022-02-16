package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/clientcredentials"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	manager     = manage.NewDefaultManager()
	srv         = server.NewServer(server.NewConfig(), manager)
	clientStore = store.NewClientStore()

	requestClient  *PublicApiInfo
	responseClient *OauthClients
	clientConfig   *clientcredentials.Config

	mariaDB *gorm.DB
	redisDB *redis.Client
)

type OauthClients struct {
	ClientID     string `gorm:"varchar(80);primary_key" json:"client_id"`
	ClientSecret string `gorm:"varchar(80);" json:"client_secret"`
	ClientIP     string `gorm:"varchar(16);" json:"client_ip"`
	GrantTypes   string `gorm:"varchar(80);" json:"grant_types"`
	Scope        string `gorm:"varchar(1600);" json:"scope"`
}

type OauthUsers struct {
	ID    string `gorm:"varchar(80);primary_key"`
	Phone string `gorm:"varchar(80)"`
	Email string `gorm:"varchar(80)"`
}

type PublicApiInfo struct {
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`
	UserID       string `json:"user_id" form:"user_id"`
}

type AuthorizationInfo struct {
	UserID      string    `json:"user_id"`
	AccessToken string    `json:"access_token"`
	Scope       []string  `json:"scope"`
	ExpiresIn   time.Time `json:"expires_in"`
}

func init() {
	initManager()
	initServer()
	initDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/oauth/token", func(c *gin.Context) {
		err := srv.HandleTokenRequest(c.Writer, c.Request)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		}
	})

	r.POST("/user/token", publicApiRequestHandler)

	r.GET("/user/info/:id", userInfoHandler)

	log.Fatal(r.Run(":9096"))
}

func initManager() {
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	manager.MustTokenStorage(store.NewMemoryTokenStore())

	manager.MapAccessGenerate(generates.NewAccessGenerate())

	manager.MapClientStorage(clientStore)
}

func initServer() {
	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})
}

func initDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mariaDB, err = gorm.Open("mysql", os.Getenv("DATA_CONNECTION_INFO"))
	if err != nil {
		log.Fatal(err.Error())
	}
	mariaDB.LogMode(true)

	redisDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func setClientStore() {
	_ = clientStore.Set(responseClient.ClientID, &models.Client{
		ID:     responseClient.ClientID,
		Secret: responseClient.ClientSecret,
		Domain: responseClient.ClientIP,
	})
}

func setClientConfig() {
	clientConfig = new(clientcredentials.Config)
	scopes := utils.SplitScope(responseClient.Scope)

	clientConfig.ClientID = responseClient.ClientID
	clientConfig.ClientSecret = responseClient.ClientSecret
	clientConfig.TokenURL = "http://localhost:9096/oauth/token"
	clientConfig.Scopes = scopes
}

func publicApiRequestHandler(c *gin.Context) {
	clientID, clientSecret := utils.BindRequestClient(c)

	if isValidClient(c) {
		setClientConfig()
		setClientStore()
	} else {
		c.JSON(500, gin.H{
			"message": "invalid client",
		})
		return
	}

	token, err := clientConfig.Token(context.Background())
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	authorizationInfo := &AuthorizationInfo{
		UserID:      requestClient.UserID,
		AccessToken: token.AccessToken,
		Scope:       clientConfig.Scopes,
		ExpiresIn:   token.Expiry,
	}
	saveAuthorizationInfo(authorizationInfo)

	c.JSON(200, authorizationInfo)
}

func userInfoHandler(c *gin.Context) {
	userID := c.Param("id")

	tokenInfo, err := srv.ValidationBearerToken(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if !isValidScope(tokenInfo.GetScope(), "write") {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "invalid scope",
		})
		return
	}

	responseUser := new(OauthUsers)
	mariaDB.Where("id = ?", userID).Find(responseUser)
	c.JSON(http.StatusOK, responseUser)
}

func saveAuthorizationInfo(authorizationInfo *AuthorizationInfo) {
	data, _ := json.Marshal(authorizationInfo)
	redisDB.Set(authorizationInfo.UserID, data, authorizationInfo.ExpiresIn.Sub(time.Now()))
}

func isValidClient(c *gin.Context) bool {
	responseClient = new(OauthClients)

	err := mariaDB.Where("client_id = ?", requestClient.ClientID).
		Where("client_secret = ?", requestClient.ClientSecret).
		Where("server_ip = ?", "1.1.1.1"). // 1.1.1.1 -> c.ClientIP()
		Find(responseClient).Error

	if err != nil {
		return false
	} else {
		return true
	}
}

func isValidScope(userScope string, apiScope string) bool {
	for _, scope := range strings.Split(userScope, " ") {
		if apiScope == scope {
			return true
		}
	}
	return false
}
