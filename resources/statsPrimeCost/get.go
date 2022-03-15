package statsPrimeCost

import (
	"github.com/gin-gonic/gin"
)

// Get more than scope 6 accessible
func Get(c *gin.Context) {
	msg := "SELECT * FROM stats_prime_cost WHERE id = 5;"
	c.Set("msg", msg)
}
