package statsPrimeCost

import (
	"github.com/gin-gonic/gin"
)

// Delete more than scope 7 accessible
func Delete(c *gin.Context) {
	msg := "원가통계 삭제"
	c.Set("msg", msg)
}
