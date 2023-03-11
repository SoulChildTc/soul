package config

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"soul/global"
	"strings"
)

func setDefaultParams(v *viper.Viper) {
	v.SetDefault("appName", "soul")
	v.SetDefault("listen", "0.0.0.0")
	v.SetDefault("port", "8080")
	v.SetDefault("env", "dev")
	//v.SetDefault("config", "app-dev.yaml") // flag中设置默认值

	// log配置
	v.SetDefault("log.path", "./app.log")
	v.SetDefault("log.level", "TRACE")
	v.SetDefault("log.console", true)
	v.SetDefault("log.closeFileLog", true)
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
	v.SetDefault("database.logLevel", "info")
	v.SetDefault("database.reportCaller", true)
	v.SetDefault("database.driver", "sqlite")
	v.SetDefault("database.path", "./data.db")

	// jwt 配置
	v.SetDefault("jwt.secret", []byte("soul"))
	v.SetDefault("jwt.ttl", "43200s") // 单位秒, 默认12小时
}
func LoadConfig() *viper.Viper {
	// 初始化viper
	v := viper.New()

	// 设置默认参数
	setDefaultParams(v)

	// 命令行参数获取
	pflag.StringP("env", "e", "dev", `运行环境, 可选项目: dev or test prod`)
	pflag.StringP("config", "c", "app-dev.yaml", `配置文件路径`)
	pflag.BoolP("migrate", "m", false, `迁移数据库`)
	pflag.Lookup("config").DefValue = "[./app-dev.yaml, ./config/app-dev.yaml]"
	pflag.Parse()

	// 命令行参数绑定
	if err := v.BindPFlags(pflag.CommandLine); err != nil {
		fmt.Println("[Init] 命令行参数绑定失败")
		panic(err.Error())
	}

	// 环境变量参数绑定
	err := v.BindEnv("env", "RUN_ENV")
	if err != nil {
		fmt.Println("[Init] 环境变量参数绑定失败")
		panic(err.Error())
	}

	filePath := v.GetString("config")

	// 不同环境读取不同配置文件
	env := strings.TrimSpace(v.GetString("env"))
	switch env {
	// 只能是下面三种环境, 如果为其他的就设置为dev
	case "dev", "test", "prod":
		if filePath != "" && v.IsSet("config") {
			// 如果指定了配置文件路径，就读取指定的配置文件
			v.SetConfigFile(filePath)
		} else {
			// 没有指定配置文件，设置对应环境的默认配置文件路径
			v.SetConfigName(fmt.Sprintf("app-%s", env))
		}
		fmt.Printf("[Init] 当前运行环境: %s\n", env)
	default:
		v.Set("env", "dev")
		v.SetConfigName("app-dev.yaml") // 默认dev环境
		fmt.Printf("[Init] 未知环境,使用默认运行环境, 默认: %s\n", "dev")
	}

	// 设置viper, 加载配置文件
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	if err := v.ReadInConfig(); err != nil {
		panic("[Init] 配置文件读取失败" + err.Error())
	}

	// 设置环境变量前缀为appName
	//v.SetEnvPrefix(v.GetString("appName"))

	// 动态加载配置
	//v.WatchConfig()
	//
	//v.OnConfigChange(func(e fsnotify.Event) {
	//	if err = v.UnmarshalExact(&global.Config); err != nil {
	//		fmt.Println("动态加载配置失败" + err.Error())
	//	}
	//})

	fmt.Printf("[Init] 使用配置文件%s\n", v.ConfigFileUsed())
	err = v.Unmarshal(&global.Config)
	if err != nil {
		panic("加载配置失败" + err.Error())
	}

	return v
}
