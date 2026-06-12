package middleware

import (
	"backend/logger"
	"backend/utils"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 获取堆栈信息
				stack := string(debug.Stack())

				// 写入日志
				logger.Errorf("panic recovered: %v\n%s", err, stack)

				// 返回错误信息
				c.JSON(utils.ServerError, utils.Response{
					Code: utils.ServerError,
					Msg:  "服务器内部错误",
				})
				// 中断请求
				c.Abort()
			}
		}()

		c.Next()
	}
}
