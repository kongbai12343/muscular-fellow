package middleware

import (
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			utils.Unauthorized(c, "身份验证失败")
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			utils.Unauthorized(c, "身份验证失败")
			c.Abort()
			return
		}

		c.Set("userId", claims.UserId)
		c.Set("userName", claims.UserName)

		c.Next()
	}
}
