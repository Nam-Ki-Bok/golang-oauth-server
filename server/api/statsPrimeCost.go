package api

import (
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/middleware"
	"infradev-practice/Wade/OAuth2.0-server/resources/statsPrimeCost"
)

// InitStatsPrimeCost handler를 클릭 시, scope를 확인할 수 있습니다.
func InitStatsPrimeCost(r *gin.Engine) {
	p := r.Group("/stats/prime-cost")
	p.Use(middleware.ValidateToken, middleware.ValidateScope, middleware.Publish)
	{
		p.GET("/:id", statsPrimeCost.Get)
		p.PUT("/:id", statsPrimeCost.Put)
		p.POST("", statsPrimeCost.Post)
		p.DELETE("/:id", statsPrimeCost.Delete)

		p.GET("", statsPrimeCost.GetList)
	}
}
