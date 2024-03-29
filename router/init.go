package router

import (
	"github.com/SoulChildTc/soul/middleware"
	"github.com/SoulChildTc/soul/router/api/v1/system"
	"github.com/SoulChildTc/soul/router/core"
	"github.com/gin-gonic/gin"
)

func registerRoute(r gin.IRouter, groupName string, register func(r *gin.RouterGroup), middlewares ...gin.HandlerFunc) {
	group := r.Group(groupName)
	group.Use(middlewares...)
	register(group)
}

func InitRouter(r *gin.Engine) {
	/*
		功能模块路由注册
	*/

	// core api
	{
		registerRoute(r, "", core.RegisterRoute)
	}

	// /api/v1
	apiV1 := r.Group("/api/v1")
	{
		registerRoute(apiV1, "/system", system.RegisterRoute)
	}

	// /api/v1 - Auth
	apiV1Auth := r.Group("/api/v1")
	apiV1Auth.Use(middleware.JwtAuth)
	{

	}

	// /api/v2
	{

	}
}
