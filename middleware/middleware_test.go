package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"infradev-practice/Wade/OAuth2.0-server/database/maria"
	"infradev-practice/Wade/OAuth2.0-server/database/redis"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	maria.Connect()
	redis.Connect()
}

func TestValidateClient(t *testing.T) {
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
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("POST", "/oauth/token", nil)
	c.Request.Header.Set("X-Real-Ip", "::1")
	c.Request.RemoteAddr = "127.0.0.1:"

	for _, tc := range cases {
		c.Request.Header.Set("Authorization", tc.auth)
		assert.Panics(t, func() { ValidateClient(c) })
	}

}

func TestValidateToken(t *testing.T) {
	
}

func TestValidateScope(t *testing.T) {

}

func TestCheckCache(t *testing.T) {

}
