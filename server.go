package main

import (
	"encoding/json"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/generates"
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/go-session/session"
)

var (
	manager = manage.NewDefaultManager()
	srv     = server.NewServer(server.NewConfig(), manager)
)

func main() {
	r := gin.Default()
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	flag.Parse()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// use redis token store
	manager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   15,
	}))

	// generate jwt access token
	//manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))
	manager.MapAccessGenerate(generates.NewAccessGenerate())

	clientStore := store.NewClientStore()
	clientStore.Set("PublicAPI", &models.Client{
		ID:     "PublicAPI",
		Secret: "test",
		Domain: "http://localhost:9094",
	})
	manager.MapClientStorage(clientStore)

	srv.SetUserAuthorizationHandler(userAuthorizeHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	r.GET("/login", loginHandler)
	r.POST("/login", loginHandler)

	r.GET("/auth", authHandler)
	r.POST("/auth", authHandler)

	r.GET("/oauth/authorize", authorizeHandler)
	r.POST("/oauth/authorize", authorizeHandler)

	r.POST("/oauth/token", func(c *gin.Context) {
		err := srv.HandleTokenRequest(c.Writer, c.Request)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		}
	})

	r.GET("/test", func(c *gin.Context) {
		token, err := srv.ValidationBearerToken(c.Request)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusBadRequest)
			return
		}

		data := map[string]interface{}{
			"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
			"client_id":  token.GetClientID(),
			"user_id":    token.GetUserID(),
			"scope":      token.GetScope(),
			"create_at":  token.GetAccessCreateAt(),
		}
		e := json.NewEncoder(c.Writer)
		e.SetIndent("", "  ")
		e.Encode(data)
	})

	log.Fatal(r.Run(":9096"))
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	log.Println("userAuthorizeHandler")

	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		store.Set("ReturnUri", r.Form)
		store.Save()

		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	userID = uid.(string)
	store.Delete("LoggedInUserID")
	store.Save()
	return
}

func loginHandler(c *gin.Context) {
	if c.Request.Method == "GET" {
		log.Println("loginHandler GET")
	}
	if c.Request.Method == "POST" {
		log.Println("loginHandler POST")
	}
	store, err := session.Start(c.Request.Context(), c.Writer, c.Request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if c.Request.Method == "POST" {
		if c.Request.Form == nil {
			if err := c.Request.ParseForm(); err != nil {
				http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		store.Set("LoggedInUserID", c.Request.Form.Get("username"))
		store.Save()

		c.Writer.Header().Set("Location", "/auth")
		c.Writer.WriteHeader(http.StatusFound)
		return
	}
	outputHTML(c, "C:\\dev\\Go\\src\\infradev-practice\\Wade\\OAuth2.0-server\\server\\static\\login.html")
}

func authHandler(c *gin.Context) {
	if c.Request.Method == "GET" {
		log.Println("authHandler GET")
	}
	if c.Request.Method == "POST" {
		log.Println("authHandler POST")
	}

	store, err := session.Start(nil, c.Writer, c.Request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		c.Writer.Header().Set("Location", "/login")
		c.Writer.WriteHeader(http.StatusFound)
		return
	}

	outputHTML(c, "C:\\dev\\Go\\src\\infradev-practice\\Wade\\OAuth2.0-server\\server\\static\\auth.html")
}

func authorizeHandler(c *gin.Context) {
	if c.Request.Method == "GET" {
		log.Println("authorizeHandler GET")
	}
	if c.Request.Method == "POST" {
		log.Println("authorizeHandler POST")
	}

	store, err := session.Start(c.Request.Context(), c.Writer, c.Request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		log.Println(v.(url.Values))
		form = v.(url.Values)
	}
	c.Request.Form = form
	store.Delete("ReturnUri")
	store.Save()

	err = srv.HandleAuthorizeRequest(c.Writer, c.Request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
	}
}

func outputHTML(c *gin.Context, filename string) {
	log.Println("outputHTML")
	file, err := os.Open(filename)
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(c.Writer, c.Request, file.Name(), fi.ModTime(), file)
}
