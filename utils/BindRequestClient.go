package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// BindClientInfo decoding basic auth value of authorization field
func BindClientInfo(c *gin.Context) (string, string) {
	authVal := c.Request.Header.Get("Authorization")
	if authVal == "" {
		ReturnError(c, http.StatusPreconditionFailed, "Header does not have authorization field")
	}

	sAuthVal := strings.Split(authVal, " ")[1]
	decAuthVal, err := base64.StdEncoding.DecodeString(sAuthVal)
	if err != nil {
		ReturnError(c, http.StatusInternalServerError, "Client id, pw decoding error")
	}

	sDecData := strings.Split(string(decAuthVal), ":")
	id := sDecData[0]
	secret := sDecData[1]

	hash := sha256.New()
	hash.Write([]byte(secret))
	shaSecret := hex.EncodeToString(hash.Sum(nil))

	return id, shaSecret
}
