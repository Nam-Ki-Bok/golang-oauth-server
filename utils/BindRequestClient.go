package utils

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// BindRequestClient decoding basic auth value of authorization
func BindRequestClient(c *gin.Context) (string, string) {
	authVal := c.Request.Header.Get("Authorization")
	if authVal == "" {
		ReturnError(c, http.StatusPreconditionFailed, "header does not have authorization field")
	}

	sAuthVal := strings.Split(authVal, " ")[1]
	decAuthVal, err := base64.StdEncoding.DecodeString(sAuthVal)
	if err != nil {
		ReturnError(c, http.StatusInternalServerError, "client id, pw decoding error")
	}

	sDecData := strings.Split(string(decAuthVal), ":")
	id := sDecData[0]
	pw := sDecData[1]

	return id, pw
}
