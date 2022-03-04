package api

import (
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/middleware"
)

// InitCost more than scope 3 accessible
func InitCost(r *gin.Engine) {
	cost := r.Group("/cost")
	cost.Use(middleware.ValidateToken, middleware.ValidateScope, middleware.Publish)
	{
		cost.GET("/:id", getCostHandler)
		cost.PUT("/:id", putCostHandler)
		cost.POST("", postCostHandler)
		cost.DELETE("/:id", delCostHandler)

		cost.GET("", getCostListHandler)
	}
}

func getCostHandler(c *gin.Context) {
	msg := "원가정보 요청"
	c.Set("msg", msg)
}

func putCostHandler(c *gin.Context) {
	msg := "원가정보 수정"
	c.Set("msg", msg)
}

func postCostHandler(c *gin.Context) {
	msg := "원가정보 삽입"
	c.Set("msg", msg)
}

func delCostHandler(c *gin.Context) {
	msg := "원가정보 삭제"
	c.Set("msg", msg)
}

func getCostListHandler(c *gin.Context) {
	msg := "원가정보 목록 요청"
	c.Set("msg", msg)
}
