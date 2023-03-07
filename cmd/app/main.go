package main

import (
	"soul/global"
	"soul/internal/config"
	"soul/internal/database"
	"soul/internal/logger"
	"soul/internal/server"
)

func init() {
	// 初始化配置文件
	global.Config = config.LoadConfig()

	// 初始化logrus
	logger.InitLogger()

	// 初始化gorm
	global.Gorm, global.SqlDB = database.InitDB()

}

func main() {
	server.StartServer()
}
