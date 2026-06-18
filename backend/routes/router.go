package routes

import (
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// 创建自定义路由
	r := gin.New()

	// 使用中间件
	r.Use(gin.Logger(), middleware.Recovery())
	// 跨域处理中间件
	r.Use(middleware.CORS())

	// 注册 404 和 405 处理
	r.NoRoute(middleware.NotFoundHandler())
	r.NoMethod(middleware.MethodNotAllowedHandler())

	return r
}
