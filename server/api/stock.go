package api

import (
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/middleware"
	"infradev-practice/Wade/OAuth2.0-server/resources/stock"
)

// InitStock handler를 클릭 시, scope를 확인할 수 있습니다.
func InitStock(r *gin.Engine) {
	s := r.Group("/stock")
	s.Use(middleware.ValidateToken, middleware.ValidateScope, middleware.Publish)
	{
		s.GET("/:id", stock.Get)
		s.PUT("/:id", stock.Put)
		s.POST("", stock.Post)
		s.DELETE("/:id", stock.Delete)

		s.GET("", stock.GetList)

	}
}
