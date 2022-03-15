package personal

import (
	"errors"
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"net/http"
)

func Post(c *gin.Context) {
	clientScope := c.GetInt("scope")
	if 1 <= clientScope && clientScope <= 3 {
		// 본인의 개인정보를 요청했는지 확인한다.
	}

	if 4 <= clientScope && clientScope <= 6 {
		// 개인정보에 접근 할 수 없는 직원 client 이다.
		utils.ReturnError(http.StatusUnauthorized, errors.New("invalid scope"))
	}

	msg := "개인 정보 삽입"
	c.Set("msg", msg)
}
