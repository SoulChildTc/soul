package config

import "github.com/spf13/viper"

func setDefaultParams(v *viper.Viper) {
	v.SetDefault("appName", "PengLai")
	v.SetDefault("listen", "0.0.0.0")
	v.SetDefault("port", "8080")
	v.SetDefault("mode", "release")
	//v.SetDefault("config", "app.yaml") // flag中设置默认值

	// log配置
	v.SetDefault("log.path", "./app.log")
	v.SetDefault("log.level", "INFO")
	v.SetDefault("log.console", true)
	v.SetDefault("log.closeFileLog", true)
	v.SetDefault("log.format", "plan")
	// log轮转
	v.SetDefault("log.rotate.enable", false)
	v.SetDefault("log.rotate.maxSize", 50)
	v.SetDefault("log.rotate.maxBackups", 0)
	v.SetDefault("log.rotate.maxAge", 7)
	v.SetDefault("log.rotate.compress", false)
	v.SetDefault("log.rotate.localtime", true)

	// database配置
	v.SetDefault("database.charset", "utf8mb4")
	v.SetDefault("database.maxOpenConns", 50)
	v.SetDefault("database.maxIdleConns", 50)
	v.SetDefault("database.connMaxIdleTime", 5)
	v.SetDefault("database.connMaxLifetime", 5)
	v.SetDefault("database.logLevel", "error")
	v.SetDefault("database.reportCaller", true)
	v.SetDefault("database.driver", "sqlite")
	v.SetDefault("database.path", "./data.db")

	// jwt 配置
	v.SetDefault("jwt.secret", []byte("penglai"))
	v.SetDefault("jwt.ttl", "43200s") // 单位秒, 默认12小时
}
