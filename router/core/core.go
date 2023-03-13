package core

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"soul/apis/controller/core"
	_ "soul/docs"
)

func RegisterRoute(r *gin.RouterGroup) {
	r.GET("/ping", core.Ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
