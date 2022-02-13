package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/clientcredentials"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	manager     = manage.NewDefaultManager()
	srv         = server.NewServer(server.NewConfig(), manager)
	clientStore = store.NewClientStore()

	requestClient  *PublicApiInfo
	responseClient *OauthClients
	clientConfig   *clientcredentials.Config

	db *gorm.DB
)

type OauthClients struct {
	ClientID     string `gorm:"varchar(80);primary_key" json:"client_id"`
	ClientSecret string `gorm:"varchar(80);" json:"client_secret"`
	ServerIP     string `gorm:"varchar(16);" json:"server_ip"`
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

	r.GET("/user/info/:id", func(c *gin.Context) {
		userID := c.Param("id")

		tokenInfo, err := srv.ValidationBearerToken(c.Request)
		log.Println(tokenInfo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		responseUser := new(OauthUsers)
		db.Where("id = ?", userID).Find(responseUser)
		c.JSON(http.StatusOK, gin.H{
			"user_id": responseUser.ID,
			"phone":   responseUser.Phone,
			"email":   responseUser.Email,
		})
	})

	log.Fatal(r.Run(":9096"))
}

func initManager() {
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// use redis token store
	manager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr: "127.0.0.1:6379",
	}))

	// generate jwt access token
	//manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))
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
	err := godotenv.Load("/Users/namkibok/KiBokFolder/Go_workspace/golang-oauth-server/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err = gorm.Open("mysql", os.Getenv("DATA_CONNECTION_INFO"))
	if err != nil {
		log.Fatal(err.Error())
	}
	db.LogMode(true)
}

func isValidClient() bool {
	responseClient = new(OauthClients)
	err := db.Where("client_id = ?", requestClient.ClientID).Find(responseClient).Error

	if err != nil || requestClient.ClientID != responseClient.ClientID || requestClient.ClientSecret != responseClient.ClientSecret {
		return false
	}

	return true
}

func setClientStore() {
	_ = clientStore.Set(responseClient.ClientID, &models.Client{
		ID:     responseClient.ClientID,
		Secret: responseClient.ClientSecret,
		Domain: responseClient.ServerIP,
	})
}

func setClientConfig() {
	clientConfig = new(clientcredentials.Config)

	clientConfig.ClientID = responseClient.ClientID
	clientConfig.ClientSecret = responseClient.ClientSecret
	clientConfig.TokenURL = "http://localhost:9096/oauth/token"
	clientConfig.Scopes = setScope()
}

func setScope() []string {
	return strings.Split(responseClient.Scope, "+")
}

func bindRequestClient(c *gin.Context) {
	requestClient = new(PublicApiInfo)
	_ = c.Bind(requestClient)
}

func publicApiRequestHandler(c *gin.Context) {
	bindRequestClient(c)

	if isValidClient() {
		setClientConfig()
		setClientStore()
	} else {
		c.JSON(500, gin.H{
			"message": "Invalid Client!",
		})
		return
	}

	token, err := clientConfig.Token(context.Background())
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(token)
	log.Println(token.TokenType)
	log.Println(token.AccessToken)
	log.Println(token.Expiry)
	c.JSON(200, gin.H{
		"user_id":      requestClient.UserID,
		"access_token": token.AccessToken,
	})
}
