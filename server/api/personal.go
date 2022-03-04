package api

import (
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/middleware"
)

// InitPersonal more than scope 4 accessible
func InitPersonal(r *gin.Engine) {
	personal := r.Group("/personal")
	personal.Use(middleware.ValidateToken, middleware.ValidateScope, middleware.Publish)
	{
		personal.GET("/:id", getPersonalHandler)
		personal.PUT("/:id", putPersonalHandler)
		personal.POST("", postPersonalHandler)
		personal.DELETE("/:id", delPersonalHandler)

		personal.GET("", getPersonalListHandler)
	}
}

func getPersonalHandler(c *gin.Context) {
	var msg string
	clientID := c.GetString("client_id")

	switch clientID {
	case "이름, 주소만 접근할 수 있는 client":
		msg = "개인정보 중 이름, 주소만 요청"
	case "성별, 나이만 접근할 수 있는 client":
		msg = "개인정보 중 성별, 나이만 요청"
	default:
		msg = "개인정보 전체 요청"
	}

	c.Set("msg", msg)
}

func putPersonalHandler(c *gin.Context) {
	msg := "개인정보 수정"
	c.Set("msg", msg)
}

func postPersonalHandler(c *gin.Context) {
	msg := "개인정보 삽입"
	c.Set("msg", msg)
}

func delPersonalHandler(c *gin.Context) {
	msg := "개인정보 삭제"
	c.Set("msg", msg)
}

func getPersonalListHandler(c *gin.Context) {
	msg := "개인정보 목록 요청"
	c.Set("msg", msg)
}
