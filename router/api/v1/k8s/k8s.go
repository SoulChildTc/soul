package k8s

import (
	"github.com/gin-gonic/gin"
	"soul/handler/controller/core"
)

// k8s模块路由

func RegisterRoute(r *gin.RouterGroup) {
	pod := r.Group("/pod")
	{
		pod.GET("/pod/detail", core.Ping)
	}

	deployment := r.Group("/deployment")
	{
		deployment.POST("/deployment/create", core.Ping)
		deployment.GET("/deployment/detail", core.Ping)
	}

}
