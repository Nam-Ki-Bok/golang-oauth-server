package personal

import (
	"github.com/gin-gonic/gin"
)

// GetList more than scope 7 accessible
func GetList(c *gin.Context) {
	msg := "개인정보 목록 요청"
	c.Set("msg", msg)
}
