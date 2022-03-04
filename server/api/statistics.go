package api

import (
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/middleware"
)

// InitStatistics more than scope 4 accessible
func InitStatistics(r *gin.Engine) {
	stat := r.Group("/statistics")
	stat.Use(middleware.ValidateToken, middleware.ValidateScope, middleware.Publish)
	{
		stat.GET("/:resource", getStatHandler)
		stat.PUT("/:resource", putStatHandler)
		stat.POST("", postStatHandler)
		stat.DELETE("/:resource", delStatHandler)

		stat.GET("", getStatListHandler)
	}
}

func getStatHandler(c *gin.Context) {
	msg := "통계정보 요청"
	c.Set("msg", msg)
}

func putStatHandler(c *gin.Context) {
	msg := "통계정보 수정"
	c.Set("msg", msg)
}

func postStatHandler(c *gin.Context) {
	msg := "통계정보 삽입"
	c.Set("msg", msg)
}

func delStatHandler(c *gin.Context) {
	msg := "통계정보 삭제"
	c.Set("msg", msg)
}

func getStatListHandler(c *gin.Context) {
	msg := "통계정보 목록 요청"
	c.Set("msg", msg)
}
