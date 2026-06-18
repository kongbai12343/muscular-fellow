package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"backend/config"
	"backend/database"
	"backend/logger"
	"backend/routes"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if err := config.Init(); err != nil {
		return err
	}

	if err := logger.Init(); err != nil {
		return err
	}
	defer logger.Sync()

	logger.Info("启动服务")

	if err := database.Init(config.Conf.Database); err != nil {
		return fmt.Errorf("数据库初始化失败: %w", err)
	}
	logger.Info("数据库初始化成功")

	router := routes.NewRouter()
	addr := fmt.Sprintf(":%d", config.Conf.Server.Port)

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	// 启动 HTTP 服务 (在 goroutine 中非阻塞运行)
	go func() {
		logger.Infof("HTTP 服务启动，监听地址: %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("服务启动失败: %v", err)
		}
	}()

	// 监听退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("收到退出信号，开始优雅关闭...")

	// 给服务 10 秒时间完成当前请求
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorf("服务关闭异常: %v", err)
		return err
	}

	logger.Info("服务已优雅关闭")
	return nil
}
