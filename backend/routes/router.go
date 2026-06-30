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
	exerciseController := controllers.NewExerciseController()

	// 路由组
	const prefix = "/api/v1"
	api := r.Group(prefix)
	{
		// 公共接口
		api.POST("/login", userController.Login)
		api.POST("/register", userController.Register)

		// 鉴权接口
		auth := api.Group("").Use(middleware.AuthMiddleware())
		{
			auth.POST("/exercises", exerciseController.Create)               // 创建动作
			auth.GET("/exercises", exerciseController.GetExercises)          // 获取所有动作
			auth.GET("/exercises/:id", exerciseController.GetExercise)       // 获取动作详情
			auth.PUT("/exercises/:id", exerciseController.UpdateExercise)    // 更新动作
			auth.DELETE("/exercises/:id", exerciseController.DeleteExercise) // 删除动作
		}
	}

	return r
}
