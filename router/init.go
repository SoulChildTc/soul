package router

import (
	"github.com/gin-gonic/gin"
	"soul/router/system"
)

func InitRouter(r *gin.Engine) {
	registerRoute := func(groupName string, f func(r *gin.RouterGroup), middlewares ...gin.HandlerFunc) {
		group := r.Group(groupName)
		for _, middle := range middlewares {
			group.Use(middle)
		}
		f(group)
	}

	registerRoute("", system.Sys)
}
