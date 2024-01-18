package system

import (
	"github.com/SoulChildTc/soul/apis/controller/system/user"
	"github.com/SoulChildTc/soul/middleware"
	"github.com/gin-gonic/gin"
)

// system模块路由

func RegisterRoute(r *gin.RouterGroup) {
	// 用户 - 免登录路由
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", user.Login)
		userGroup.POST("/register", user.Register)
	}

	// 用户 - 需JWT路由
	userAuthGroup := userGroup.Group("").Use(middleware.JwtAuth)
	{
		userAuthGroup.GET("/info", user.Info)
	}
}
