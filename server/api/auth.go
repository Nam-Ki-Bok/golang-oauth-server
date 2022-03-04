package api

import (
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/middleware"
	"infradev-practice/Wade/OAuth2.0-server/token"
)

func InitAuth(r *gin.Engine) {
	auth := r.Group("/oauth")
	auth.Use(middleware.ValidateClient, middleware.CheckCache)
	auth.POST("/token", token.Request)

	// set token endpoint
	r.POST("/token", token.Generate)
}
