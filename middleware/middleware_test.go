package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"infradev-practice/Wade/OAuth2.0-server/database/maria"
	"infradev-practice/Wade/OAuth2.0-server/database/redis"
	model "infradev-practice/Wade/OAuth2.0-server/models"
	"infradev-practice/Wade/OAuth2.0-server/token"
)

var (
	w    = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	maria.Connect()
	redis.Connect()

	c.Request, _ = http.NewRequest("", "", nil)
	c.Request.Header.Set("X-Real-Ip", "::1")
	c.Request.RemoteAddr = "127.0.0.1:"

	token.Cs.Set("test", &models.Client{
		ID:     "test",
		Secret: "test",
	})
}

func TestSuccessValidateClient(t *testing.T) {
	assert.New(t)

	// given
	c.Request.Header.Set("Authorization", "Basic dGVzdDp0ZXN0")

	// when
	ValidateClient(c)

	// then
	assert.Equal(t, "test", c.GetString("client_id"))
	assert.Equal(t, "9F86D081884C7D659A2FEAA0C55AD015A3BF4F1B2B0B822CD15D6C15B0F00A08",
		c.GetString("client_secret"))
}

func TestPanicValidateClient(t *testing.T) {
	assert.New(t)
	type testCase struct {
		auth string
	}

	cases := []testCase{
		{auth: "Basic Zm9vOmJhcg=="}, // foo:bar
		{auth: "Basic Zm8gbzpiYXI="}, // fo o:bar
		{auth: "Basic Zm9vOmJhIHI="}, // foo:ba r
		{auth: "Basic Zm9vOg=="},     // foo:
		{auth: "Basic Zm9vOmJhIHI="}, // foo:ba r
		{auth: "Basic OmJhcg=="},     // :bar
		{auth: "no auth"},
	}

	for _, tc := range cases {
		c.Request.Header.Set("Authorization", tc.auth)
		assert.Panics(t, func() { ValidateClient(c) })
	}
}

func TestSuccessValidateToken(t *testing.T) {
	assert.New(t)

	ti, _ := token.Mgr.GenerateAccessToken(context.Background(), oauth2.ClientCredentials, &oauth2.TokenGenerateRequest{
		ClientID:     "test",
		ClientSecret: "test",
		Scope:        "0",
	})
	c.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ti.GetAccess()))

	ValidateToken(c)

	assert.Equal(t, "test", c.GetString("client_id"))
	assert.Equal(t, 0, c.GetInt("scope"))
}

func TestPanicValidateToken(t *testing.T) {
	assert.New(t)

	ti, _ := token.Mgr.GenerateAccessToken(context.Background(), oauth2.ClientCredentials, &oauth2.TokenGenerateRequest{
		ClientID:     "test",
		ClientSecret: "test",
		Scope:        "no scope",
	})
	c.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ti.GetAccess()))
	assert.Panics(t, func() { ValidateToken(c) })

	c.Request.Header.Set("Authorization", "Bearer no")
	assert.Panics(t, func() { ValidateToken(c) })
}

func TestPanicValidateScope(t *testing.T) {
	assert.New(t)

	type testCase struct {
		url string
	}

	cases := []testCase{
		{url: "/foo"},
		{url: "/foo/bar"},
		{url: "/foo/bar/1"},
		{url: "/test"},
	}

	for _, tc := range cases {
		c.Set("scope", 0)
		c.Request, _ = http.NewRequest("GET", tc.url, nil)
		assert.Panics(t, func() { ValidateScope(c) })
	}
}

func TestCheckCache(t *testing.T) {
	var output map[string]string

	cache := model.AuthInfo{
		ClientID:    "test",
		AccessToken: "token",
	}
	data, _ := json.Marshal(cache)
	redis.DB.Set(context.Background(), cache.ClientID, data, 10)

	c.Request, _ = http.NewRequest("", "", nil)
	c.Set("client_id", "test")

	CheckCache(c)

	_ = json.Unmarshal(w.Body.Bytes(), &output)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "test", output["client_id"])
	assert.Equal(t, "token", output["access_token"])

	redis.DB.Del(context.Background(), "test")
}
