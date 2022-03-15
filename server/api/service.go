package api

import (
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/resources/service"
)

// InitService handler를 클릭 시, scope를 확인할 수 있습니다.
func InitService(r *gin.Engine) {
	s := r.Group("/service")
	{
		s.GET("/:id", service.Get)
		s.PUT("/:id", service.Put)
		s.POST("", service.Post)
		s.DELETE("/:id", service.Delete)

		s.GET("", service.GetList)
	}
}
