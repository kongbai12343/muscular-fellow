package routes

import (
	"backend/controllers"
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

	// 初始化controller

	userController := controllers.NewUserController()

	// 路由组
	const prefix = "/api/v1"
	api := r.Group(prefix)
	{
		api.POST("/login", userController.Login)
		api.POST("/register", userController.Register)
	}

	return r
}
