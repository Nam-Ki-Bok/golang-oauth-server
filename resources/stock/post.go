package stock

import (
	"github.com/gin-gonic/gin"
)

// Post more than scope 4 accessible
func Post(c *gin.Context) {
	msg := "재고정보 삽입"
	c.Set("msg", msg)
}
