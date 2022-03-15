package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/database/maria"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"net/http"
	"regexp"
)

type apiMaps struct {
	Resource string
	Method   string
	Id       bool
	Scope    int
}

func ValidateScope(c *gin.Context) {
	method := c.Request.Method
	resource, isExistID := parseResource(c.Request.URL.Path)

	api := new(apiMaps)
	result := maria.DB.Where("resource = ?", resource).
		Where("method = ?", method).
		Where("id = ?", isExistID).
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

func parseResource(uri string) (string, bool) {
	isExistID, _ := regexp.MatchString("/([0-9]+)", uri)

	r := regexp.MustCompile("/([0-9]+)")
	resource := r.ReplaceAllString(uri, "")

	return resource, isExistID
}
