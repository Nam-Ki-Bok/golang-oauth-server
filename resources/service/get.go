package service

import (
	"github.com/gin-gonic/gin"
)

// Get more than scope 3 accessible
// client scope == 3 인 경우, 본인의 프로젝트인지 확인해야 한다.
// client scope >= 4 인 경우, 직원이기 때문에 본인의 프로젝트인지 확인할 필요가 없다.
func Get(c *gin.Context) {
	clientScope := c.GetInt("scope")
	if clientScope == 3 {
		// 본인의 프로젝트인지 확인해야 한다.
	}

	// 직원이기 때문에 본인의 프로젝트인지 확인할 필요가 없다.
}
