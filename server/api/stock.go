package api

import (
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/middleware"
)

// InitStock more than scope 2 accessible
func InitStock(r *gin.Engine) {
	stock := r.Group("/stock")
	stock.Use(middleware.ValidateToken, middleware.ValidateScope, middleware.Publish)
	{
		stock.GET("/:id", getStockHandler)
		stock.PUT("/:id", putStockHandler)
		stock.POST("", postStockHandler)
		stock.DELETE("/:id", delStockHandler)

		stock.GET("", getStockListHandler)
	}
}

func getStockHandler(c *gin.Context) {
	msg := "재고정보 요청"
	c.Set("msg", msg)
}

func putStockHandler(c *gin.Context) {
	msg := "재고정보 수정"
	c.Set("msg", msg)
}

func postStockHandler(c *gin.Context) {
	msg := "재고정보 삽입"
	c.Set("msg", msg)
}

func delStockHandler(c *gin.Context) {
	msg := "재고정보 삭제"
	c.Set("msg", msg)
}

func getStockListHandler(c *gin.Context) {
	msg := "재고정보 목록 요청"
	c.Set("msg", msg)
}
