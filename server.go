package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
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
)

var (
	manager     = manage.NewDefaultManager()
	srv         = server.NewServer(server.NewConfig(), manager)
	clientStore = store.NewClientStore()

	db *gorm.DB
)

type OauthClients struct {
	ClientID     string `gorm:"varchar(80);primary_key" json:"client_id"`
	ClientSecret string `gorm:"varchar(80);primary_key" json:"client_secret"`
	ServerIP     string `gorm:"varchar(16);primary_key" json:"server_ip"`
	GrantTypes   string `gorm:"varchar(80);primary_key" json:"grant_types"`
	Scope        string `gorm:"varchar(1600);primary_key" json:"scope"`
}

func init() {
	initManager()
	initServer()
	initDatabase()
}

func isValidClient() {

}

func setClientStore() {

}

func setScope() {

}

func main() {
	r := gin.Default()

	r.POST("/oauth/token", func(c *gin.Context) {
		err := srv.HandleTokenRequest(c.Writer, c.Request)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		}
	})

	r.GET("/user/token", func(c *gin.Context) {
		requestDTO := new(clientcredentials.Config)
		_ = c.Bind(requestDTO)

		token, err := requestDTO.Token(context.Background())
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println(token)
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

	//clientStore.Set("PublicAPI", &models.Client{
	//	ID:     "PublicAPI",
	//	Secret: "test",
	//	Domain: "http://localhost:9094",
	//})
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
