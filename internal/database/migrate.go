package database

import (
	"fmt"
	"soul/global"
)

func InitDBMigrate() {
	err := global.DB.AutoMigrate(
	//your model. eg: model.ID{},
	)

	if err != nil {
		panic("迁移数据库模型失败! " + err.Error())
	}
	fmt.Println("[Init] 数据库模型迁移成功")
}
