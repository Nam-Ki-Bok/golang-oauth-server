package statsPrimeCost

import (
	"github.com/gin-gonic/gin"
)

// Put more than scope 7 accessible
func Put(c *gin.Context) {
	msg := "원가통계 수정"
	c.Set("msg", msg)
}
