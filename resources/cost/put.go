package cost

import (
	"github.com/gin-gonic/gin"
)

// Put more than scope 2 accessible
// 2 <= client scope <= 3 인 경우, 본인의 프로젝트인지 확인해야 한다.
// client scope >= 4 인 경우, 직원이기 때문에 본인의 프로젝트인지 확인할 필요가 없다.
func Put(c *gin.Context) {
	clientScope := c.GetInt("scope")
	if 2 <= clientScope && clientScope <= 3 {
		// 본인의 프로젝트인지 확인해야 한다.
	}

	// 직원이기 때문에 본인의 프로젝트인지 확인할 필요가 없다.
	msg := "요금 정보 수정"
	c.Set("msg", msg)
}
