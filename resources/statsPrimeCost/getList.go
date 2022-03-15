package statsPrimeCost

import (
	"github.com/gin-gonic/gin"
)

// GetList more than scope 6 accessible
func GetList(c *gin.Context) {
	msg := "SELECT * FROM stats_prime_cost;"
	c.Set("msg", msg)
}
