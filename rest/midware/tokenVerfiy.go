package midware

import (
	"dtyTrade/rest/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IjU5NzAzMTIyQHFxLmNvbSIsImV4cGlyZWRBdCI6MTU3MTkwNDEwMiwiaWQiOjJ9.nRYdO5shFhVQ3VBzUunsftUMKF0xD0jQXzyGpky6zlk
func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if len(token) == 0 {
			token = c.Query("token")
			if len(token) == 0 {
				c.AbortWithStatusJSON(http.StatusForbidden, "无效的token")
				return
			}
		}
		// 在gin上下文中定义变量
		user, err := service.CheckToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, "无效的token")
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
