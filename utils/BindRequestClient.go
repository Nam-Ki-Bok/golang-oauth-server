package utils

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// BindRequestClient decoding basic auth value of authorization
func BindRequestClient(c *gin.Context) (id, pw string, err error) {
	data := strings.Split(c.Request.Header.Get("Authorization"), " ")[1]

	sDec, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		ReturnError(c, http.StatusInternalServerError, "client id, pw decoding error")
	}

	split := strings.Split(string(sDec), ":")
	id = split[0]
	pw = split[1]

	return id, pw, nil
}
