package stock

import (
	"github.com/gin-gonic/gin"
)

// Delete more than scope 4 accessible
func Delete(c *gin.Context) {
	msg := "재고정보 삭제"
	c.Set("msg", msg)
}
