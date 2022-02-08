package main

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	authServerURL = "http://localhost:9096"
)

var (
	config = oauth2.Config{
		ClientID:     "PublicAPI",
		ClientSecret: "test",
		Scopes:       []string{"client"},
		RedirectURL:  "http://localhost:9094/oauth2",
		Endpoint: oauth2.Endpoint{
			AuthURL:  authServerURL + "/oauth/authorize",
			TokenURL: authServerURL + "/oauth/token",
		},
	}
	globalToken *oauth2.Token
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		log.Println("GET /")

		u := config.AuthCodeURL("wade",
			oauth2.SetAuthURLParam("code_challenge", genCodeChallengeS256("s256example")),
			oauth2.SetAuthURLParam("code_challenge_method", "S256"))
		log.Printf("Redirect URL : %s", u)
		c.Redirect(http.StatusFound, u)
	})

	r.GET("/oauth2", func(c *gin.Context) {
		log.Println("client /oauth2")

		_ = c.Request.ParseForm()
		state := c.Request.Form.Get("state")
		if state != "wade" {
			http.Error(c.Writer, "State invalid", http.StatusBadRequest)
			return
		}

		// Authorization Code 추출
		code := c.Request.Form.Get("code")
		if code == "" {
			http.Error(c.Writer, "Code not found", http.StatusBadRequest)
			return
		}

		// 토큰 발급 과정
		token, err := config.Exchange(context.Background(), code, oauth2.SetAuthURLParam("code_verifier", "s256example"))
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
		globalToken = token

		e := json.NewEncoder(c.Writer)
		e.SetIndent("", "  ")
		e.Encode(token)
	})

	r.GET("/refresh", func(c *gin.Context) {
		log.Println("client /refresh")

		if globalToken == nil {
			http.Redirect(c.Writer, c.Request, "/", http.StatusFound)
			return
		}

		globalToken.Expiry = time.Now()
		token, err := config.TokenSource(context.Background(), globalToken).Token()
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		globalToken = token
		e := json.NewEncoder(c.Writer)
		e.SetIndent("", "  ")
		e.Encode(token)
	})

	r.GET("/try", func(c *gin.Context) {
		log.Println("client /try")

		if globalToken == nil {
			c.Redirect(http.StatusFound, "/")
			return
		}

		// Resource Server에 토큰과 함께 정보를 요청
		resp, err := http.Get(fmt.Sprintf("%s/test?access_token=%s", authServerURL, globalToken.AccessToken))
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusBadRequest)
			return
		}
		defer resp.Body.Close()

		io.Copy(c.Writer, resp.Body)
	})

	log.Println("Client is running at 9094 port.Please open http://localhost:9094")
	log.Fatal(r.Run(":9094"))
}

func genCodeChallengeS256(s string) string {
	s256 := sha256.Sum256([]byte(s))
	return base64.URLEncoding.EncodeToString(s256[:])
}
