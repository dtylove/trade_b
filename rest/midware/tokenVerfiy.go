package midware

import (
	"dtyTrade/rest/response"
	"dtyTrade/rest/service"
	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if len(token) == 0 {
			token = c.Query("token")
			if len(token) == 0 {
				response.Res(c, response.C_TOKEN_NOT_FOUND, "token is null")
				return
			}
		}
		// 在gin上下文中定义变量
		user, err := service.CheckToken(token)
		if err != nil {
			response.Res(c, response.C_TOKEN_NOT_FOUND, "token not found")
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
