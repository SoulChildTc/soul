package k8s

import (
	"github.com/SoulChildTc/soul/apis/controller/core"
	"github.com/gin-gonic/gin"
)

// k8s模块路由

func RegisterRoute(r *gin.RouterGroup) {
	r.GET("/deployment", core.Ping)
}
