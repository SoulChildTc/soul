package middleware

import (
	"github.com/gin-gonic/gin"
)

// InitMiddleware 加载全局中间件
func InitMiddleware(r *gin.Engine) {
	// panic 拦截器
	r.Use(ErrorInterceptor)
	// 日志处理
	r.Use(Logger)

	// 设置 X-Request-Id header
	r.Use(RequestId)

	// JWT Auth
	//r.Use(JwtAuth)

	// 跨域处理
	r.Use(Cors)
}
