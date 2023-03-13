package system

import (
	"github.com/gin-gonic/gin"
	"soul/apis/controller/system/user"
)

// system模块路由

func RegisterRoute(r *gin.RouterGroup) {
	// 用户
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", user.Login)
		userGroup.POST("/register", user.Register)
	}
}
