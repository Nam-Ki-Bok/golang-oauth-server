package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/generates"
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/go-redis/redis/v8"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
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
	dumpvar   bool
	idvar     string
	secretvar string
	domainvar string
	portvar   int

	manager = manage.NewDefaultManager()
	srv     = server.NewServer(server.NewConfig(), manager)
)

func init() {
	flag.BoolVar(&dumpvar, "d", true, "Dump requests and responses")
	flag.StringVar(&idvar, "i", "PublicAPI", "The client id being passed in")
	flag.StringVar(&secretvar, "s", "test", "The client secret being passed in")
	flag.StringVar(&domainvar, "r", "http://localhost:9094", "The domain of the redirect url")
	flag.IntVar(&portvar, "p", 9096, "the base port for the server")
}

func main() {
	r := gin.Default()

	flag.Parse()
	if dumpvar {
		log.Println("Dumping requests")
	}

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
	clientStore.Set(idvar, &models.Client{
		ID:     idvar,
		Secret: secretvar,
		Domain: domainvar,
	})
	manager.MapClientStorage(clientStore)

	srv.SetPasswordAuthorizationHandler(func(ctx context.Context, username, password string) (userID string, err error) {
		if username == "test" && password == "test" {
			userID = "test"
		}
		return
	})

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
		if dumpvar {
			_ = dumpRequest(os.Stdout, "token", c.Request) // Ignore the error
		}

		err := srv.HandleTokenRequest(c.Writer, c.Request)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		}
	})

	r.GET("/test", func(c *gin.Context) {
		if dumpvar {
			_ = dumpRequest(os.Stdout, "test", c.Request) // Ignore the error
		}
		token, err := srv.ValidationBearerToken(c.Request)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusBadRequest)
			return
		}

		data := map[string]interface{}{
			"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
			"client_id":  token.GetClientID(),
			"user_id":    token.GetUserID(),
		}
		e := json.NewEncoder(c.Writer)
		e.SetIndent("", "  ")
		e.Encode(data)
	})

	log.Printf("Server is running at %d port.\n", portvar)
	log.Printf("Point your OAuth client Auth endpoint to %s:%d%s", "http://localhost", portvar, "/oauth/authorize")
	log.Printf("Point your OAuth client Token endpoint to %s:%d%s", "http://localhost", portvar, "/oauth/token")
	log.Fatal(r.Run(fmt.Sprintf(":%d", portvar)))
}

func dumpRequest(writer io.Writer, header string, r *http.Request) error {
	data, err := httputil.DumpRequest(r, true)
	if err != nil {
		return err
	}
	writer.Write([]byte("\n" + header + ": \n"))
	writer.Write(data)
	return nil
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	if dumpvar {
		_ = dumpRequest(os.Stdout, "userAuthorizeHandler", r) // Ignore the error
	}
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
	if dumpvar {
		_ = dumpRequest(os.Stdout, "login", c.Request) // Ignore the error
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

		// /auth로 리다이렉트
		c.Writer.Header().Set("Location", "/auth")
		c.Writer.WriteHeader(http.StatusFound)
		return
	}
	outputHTML(c, "C:\\dev\\Go\\src\\infradev-practice\\Wade\\OAuth2.0-server\\server\\static\\login.html")
}

func authHandler(c *gin.Context) {
	if dumpvar {
		_ = dumpRequest(os.Stdout, "auth", c.Request) // Ignore the error
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
	if dumpvar {
		dumpRequest(os.Stdout, "authorize", c.Request)
	}

	store, err := session.Start(c.Request.Context(), c.Writer, c.Request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
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
	file, err := os.Open(filename)
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(c.Writer, c.Request, file.Name(), fi.ModTime(), file)
}
