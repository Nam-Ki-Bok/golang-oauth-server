package statsPrimeCost

import (
	"github.com/gin-gonic/gin"
)

// Get more than scope 6 accessible
func Get(c *gin.Context) {
	msg := "원가통계 요청"
	c.Set("msg", msg)
}
