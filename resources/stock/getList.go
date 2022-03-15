package stock

import (
	"github.com/gin-gonic/gin"
)

// GetList more than scope 4 accessible
func GetList(c *gin.Context) {
	msg := "재고정보 목록 요청"
	c.Set("msg", msg)
}
