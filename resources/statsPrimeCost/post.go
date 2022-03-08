package statsPrimeCost

import (
	"github.com/gin-gonic/gin"
)

// Post more than scope 7 accessible
func Post(c *gin.Context) {
	msg := "원가통계 삽입"
	c.Set("msg", msg)
}
