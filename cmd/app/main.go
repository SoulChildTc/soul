package cmd

import (
	"github.com/SoulChildTc/soul/global"
	"github.com/SoulChildTc/soul/internal/config"
	"github.com/SoulChildTc/soul/internal/database"
	"github.com/SoulChildTc/soul/internal/logger"
	"github.com/SoulChildTc/soul/internal/server"
	"os"
)

func init() {
	//初始化配置文件
	global.V = config.LoadConfig()

	// 初始化logrus
	logger.InitLogger()

	// 初始化gorm
	global.DB, global.SqlDB = database.InitDB()

	// 数据迁移
	if global.V.GetBool("migrate") {
		database.InitDBMigrate()
		os.Exit(0)
	}

}

func Execute() {
	server.StartServer()
}
