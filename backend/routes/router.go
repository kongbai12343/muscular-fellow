package routes

import (
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	// 创建自定义路由
	r := gin.New()

	// 使用中间件
	r.Use(gin.Logger(), middleware.Recovery())

	return r
}
