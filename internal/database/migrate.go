package database

import (
	"fmt"
	"github.com/SoulChildTc/soul/global"
	"github.com/SoulChildTc/soul/model"
)

func InitDBMigrate() {
	err := global.DB.AutoMigrate(model.MigrateModels...)

	if err != nil {
		panic("迁移数据库模型失败! " + err.Error())
	}
	fmt.Println("[Init] 数据库模型迁移成功")
}
