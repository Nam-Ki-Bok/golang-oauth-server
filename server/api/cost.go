package api

import (
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/middleware"
	"infradev-practice/Wade/OAuth2.0-server/resources/cost"
)

// InitCost handler를 클릭 시, scope를 확인할 수 있습니다.
func InitCost(r *gin.Engine) {
	c := r.Group("/cost")
	c.Use(middleware.ValidateToken, middleware.ValidateScope, middleware.Publish)
	{
		c.GET("/:id", cost.Get)
		c.PUT("/:id", cost.Put)
		c.POST("", cost.Post)
		c.DELETE("/:id", cost.Delete)

		c.GET("", cost.GetList)
	}
}
