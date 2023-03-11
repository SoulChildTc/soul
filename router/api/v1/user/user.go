package user

import (
	"github.com/gin-gonic/gin"
	"soul/handler/controller/core"
)

// 用户模块路由

func RegisterRoute(r *gin.RouterGroup) {
	r.Group("/user")
	{
		r.GET("/login", core.Ping)
	}
}
