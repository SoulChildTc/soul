package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"soul/global"
	log "soul/internal/logger"
	"time"
)

func InitDB() (*gorm.DB, *sql.DB) {
	dbDriver := global.Config.GetString("database.driver")

	switch dbDriver {
	case "mysql":
		fmt.Printf("[Init] 使用%s数据库驱动\n", dbDriver)
		return initMySqlGorm()
	case "sqlite":
		//return initSqliteGorm()
		fmt.Printf("[Init] 使用%s数据库驱动\n", dbDriver)
		return nil, nil
	default:
		//return initSqliteGorm()
		fmt.Println("[Init] 使用默认数据库驱动 - Sqlite")
		return nil, nil
	}
}

func initMySqlGorm() (*gorm.DB, *sql.DB) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		global.Config.GetString("database.username"),
		global.Config.GetString("database.password"),
		global.Config.GetString("database.host"),
		global.Config.GetInt("database.port"),
		global.Config.GetString("database.database"),
		global.Config.GetString("database.charset"),
	)

	mysqlConfig := mysql.Config{DSN: dsn}

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
		Logger: log.NewGormLogger(),
	}

	gormDB, err := gorm.Open(mysql.New(mysqlConfig), gormConfig)
	if err != nil {
		os.Exit(1)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		panic("数据库连接池初始化失败")
	}
	sqlDB.SetMaxIdleConns(global.Config.GetInt("database.MaxIdleConns"))
	sqlDB.SetMaxOpenConns(global.Config.GetInt("database.MaxOpenConns"))

	connMaxIdleTime := time.Duration(global.Config.GetInt("database.MaxOpenConns")) * time.Minute
	connMaxLifetime := time.Duration(global.Config.GetInt("database.MaxOpenConns")) * time.Minute
	sqlDB.SetConnMaxIdleTime(connMaxIdleTime)
	sqlDB.SetConnMaxLifetime(connMaxLifetime)

	log.Info("数据库连接初始化成功")
	return gormDB, sqlDB
}

func initSqliteGorm() (*gorm.DB, *sql.DB) {
	return nil, nil
}
