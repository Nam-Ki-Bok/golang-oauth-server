package stock

import (
	"github.com/gin-gonic/gin"
)

// Get more than scope 4 accessible
func Get(c *gin.Context) {
	msg := "재고정보 요청"
	c.Set("msg", msg)
}
