package system

import (
	"github.com/gin-gonic/gin"
	"soul/controller"
)

func Sys(r *gin.RouterGroup) {
	r.GET("/ping", controller.Ping)
	r.GET("/ping2", controller.Ping)
}
