package core

import (
	"github.com/SoulChildTc/soul/apis/controller/core"
	_ "github.com/SoulChildTc/soul/docs"
	"github.com/SoulChildTc/soul/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoute(r *gin.RouterGroup) {
	r.GET("/ping", core.Ping)
	swag := r.Group("/swagger").Use(middleware.BasicAuth)
	swag.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
