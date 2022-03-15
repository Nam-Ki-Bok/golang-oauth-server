package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/database/redis"
	"infradev-practice/Wade/OAuth2.0-server/models"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"net/http"
)

// CheckCache If there is an auth record,
// Auth record will be returned without going through the token generation process
func CheckCache(c *gin.Context) {
	id := c.GetString("client_id")
	authInfo := new(models.AuthInfo)
	cache := redis.DB.Get(c, id)

	switch {
	case cache.Err() == redis.Nil:
		c.Next()
	case cache.Err() != redis.Nil:
		err := json.Unmarshal([]byte(cache.Val()), &authInfo)
		if err != nil {
			utils.ReturnError(http.StatusBadRequest, err)
		}
		c.AbortWithStatusJSON(http.StatusOK, authInfo)
	}
}
