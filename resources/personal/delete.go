package personal

import (
	"errors"
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"net/http"
)

// Delete more than scope 1 accessible
// scope <= 3 인 경우, 본인인지 확인 하는 과정이 필요하다.
//  ex) id = 3 인 사람이 id = 2 의 정보를 GET, PUT, DELETE 하면 안된다.
func Delete(c *gin.Context) {
	clientScope := c.GetInt("scope")
	if 1 <= clientScope && clientScope <= 3 {
		// 본인의 개인정보를 요청했는지 확인한다.
	}

	if 4 <= clientScope && clientScope <= 6 {
		// 개인정보에 접근 할 수 없는 직원 client 이다.
		utils.ReturnError(http.StatusUnauthorized, errors.New("invalid scope"))
	}

	msg := "개인 정보 삭제"
	c.Set("msg", msg)
}
