package api

import (
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/middleware"
	"infradev-practice/Wade/OAuth2.0-server/resources/personal"
)

// InitPersonal handler를 클릭 시, scope를 확인할 수 있습니다.
func InitPersonal(r *gin.Engine) {
	p := r.Group("/personal")
	p.Use(middleware.ValidateToken, middleware.ValidateScope, middleware.Publish)
	{
		p.GET("/:id", personal.Get)
		p.PUT("/:id", personal.Put)
		p.POST("", personal.Post)
		p.DELETE("/:id", personal.Delete)

		p.GET("", personal.GetList)
	}
}
