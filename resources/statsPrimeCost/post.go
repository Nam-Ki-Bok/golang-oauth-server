package statsPrimeCost

import (
	"github.com/gin-gonic/gin"
)

// Post more than scope 7 accessible
func Post(c *gin.Context) {
	msg := "INSERT INTO table VALUES (\"foo\", \"bar\");"
	c.Set("msg", msg)
}
