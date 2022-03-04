package api

import (
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/middleware"
)

// InitOwn more than scope 1 accessible
func InitOwn(r *gin.Engine) {
	own := r.Group("/own")
	own.Use(middleware.ValidateToken, middleware.ValidateScope, middleware.Publish)
	{
		own.GET("/:id", getOwnHandler)
		own.PUT("/:id", putOwnHandler)
		own.POST("", postOwnHandler)
		own.DELETE("/:id", delOwnHandler)

		own.GET("", getOwnListHandler)
	}
}

func getOwnHandler(c *gin.Context) {
	msg := "본인정보 요청"
	c.Set("msg", msg)
}

func putOwnHandler(c *gin.Context) {
	msg := "본인정보 수정"
	c.Set("msg", msg)
}

func postOwnHandler(c *gin.Context) {
	msg := "본인정보 삽입"
	c.Set("msg", msg)
}

func delOwnHandler(c *gin.Context) {
	msg := "본인정보 삭제"
	c.Set("msg", msg)
}

func getOwnListHandler(c *gin.Context) {
	msg := "본인정보 목록 요청"
	c.Set("msg", msg)
}
