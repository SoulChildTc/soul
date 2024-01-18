package server

import (
	"context"
	"fmt"
	"github.com/SoulChildTc/soul/global"
	"github.com/SoulChildTc/soul/internal/logger"
	"github.com/SoulChildTc/soul/middleware"
	"github.com/SoulChildTc/soul/router"
	"github.com/SoulChildTc/soul/utils/httputil"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartServer() {
	switch global.Config.Mode {
	case "test":
		gin.SetMode(gin.TestMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// 初始化全局中间件
	middleware.InitMiddleware(r)

	// 初始化路由和路由中间件
	router.InitRouter(r)

	// 注册自定义验证器
	httputil.RegisterAllValidator()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", global.Config.Listen, global.Config.Port),
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // kill -2 和 -15

	// 等待终止信号
	<-quit

	logger.Info("Shutdown Server ...")

	// 创建一个5秒超时的ctx
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅停止, 如果有正在处理的请求,等待ctx超时或cancel后强制停止, 同时停止接收新的请求
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}
	logger.Info("Server exiting")
}
