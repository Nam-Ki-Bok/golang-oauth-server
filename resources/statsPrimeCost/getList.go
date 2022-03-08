package statsPrimeCost

import (
	"github.com/gin-gonic/gin"
)

// GetList more than scope 6 accessible
func GetList(c *gin.Context) {
	msg := "원가통계 목록 요청"
	c.Set("msg", msg)
}
