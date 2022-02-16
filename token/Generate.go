package token

import (
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"net/http"
)

func Generate(c *gin.Context) {
	err := srv.HandleTokenRequest(c.Writer, c.Request)
	if err != nil {
		utils.ReturnError(c, http.StatusInternalServerError, "Failed to generate a Token")
	}
}
