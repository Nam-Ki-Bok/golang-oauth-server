package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/token"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"net/http"
	"strconv"
)

func ValidateToken(c *gin.Context) {
	// Token.Srv uses the token authentication process of go-oauth2.
	tokenInfo, err := token.Srv.ValidationBearerToken(c.Request)
	if err != nil {
		utils.ReturnError(http.StatusUnauthorized, err)
	}

	scope, err := strconv.Atoi(tokenInfo.GetScope())
	if err != nil {
		utils.ReturnError(http.StatusBadRequest, errors.New("failed to convert string scope to integer scope"))
	}

	c.Set("client_id", tokenInfo.GetClientID())
	c.Set("scope", scope)
	c.Next()
}
