package middleware

import (
	"backend/logger"
	"backend/utils"
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stack := string(debug.Stack())

				// 写入日志
				logger.Errorf("panic recovered: %v\n%s", err, stack)

				utils.ServerError(c, "服务内部错误")
				c.Abort()
			}
		}()

		c.Next()
	}
}

func NotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.NotFound(c, fmt.Sprintf("接口不存在: %s %s", c.Request.Method, c.Request.URL.Path))
	}
}

func MethodNotAllowedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.MethodNotAllowed(c, fmt.Sprintf("请求方法不允许: %s %s", c.Request.Method, c.Request.URL.Path))
	}
}
