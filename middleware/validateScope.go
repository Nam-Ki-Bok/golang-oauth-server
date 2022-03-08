package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/database/maria"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"net/http"
	"strconv"
	"strings"
)

type apiMaps struct {
	Resource string
	Method   string
	Id       bool
	Scope    int
}

func ValidateScope(c *gin.Context) {
	method := c.Request.Method
	resource, isID := parseResource(c.Request.URL.Path)

	api := new(apiMaps)
	result := maria.DB.Where("resource = ?", resource).
		Where("method = ?", method).
		Where("id = ?", isID).
		Find(&api)
	if result.RowsAffected == 0 {
		utils.ReturnError(http.StatusBadRequest, errors.New("invalid request"))
	}

	scope := c.GetInt("scope")

	if scope < api.Scope {
		utils.ReturnError(http.StatusUnauthorized, errors.New("invalid scope"))
	}

	c.Next()
}

// parseResource
// "/resource len = 2"
// "/resource/:id len = 3"
// "/stats/resource len = 3"
// "/stats/resource/:id len = 4"
func parseResource(uri string) (resource string, isID bool) {
	splitURI := strings.Split(uri, "/")
	resource = splitURI[1]

	switch len(splitURI) {
	case 2:
		isID = false
	case 3:
		isID = true
		if _, err := strconv.Atoi(splitURI[2]); err != nil {
			resource = resource + "-" + splitURI[2]
		}
	case 4:
		resource = resource + "-" + splitURI[2]
		isID = true
	}
	return
}
