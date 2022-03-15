package stock

import (
	"github.com/gin-gonic/gin"
)

// Put more than scope 4 accessible
func Put(c *gin.Context) {
	msg := "재고정보 수정"
	c.Set("msg", msg)
}
