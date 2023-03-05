package config

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
)

var v *viper.Viper

func setDefaultParams(v *viper.Viper) {
	v.SetDefault("port", "8080")
	v.SetDefault("env", "dev")

	// log配置
	v.SetDefault("log.path", "./app.log")
	v.SetDefault("log.rotate.enable", false)
	v.SetDefault("log.rotate.maxSize", 50)
	v.SetDefault("log.rotate.maxBackups", 0)
	v.SetDefault("log.rotate.maxAge", 7)
	v.SetDefault("log.rotate.compress", false)
	v.SetDefault("log.rotate.localtime", true)
}

func LoadConfig() {
	// 初始化viper
	v = viper.New()

	// 设置默认参数
	setDefaultParams(v)

	// 命令行参数获取
	pflag.StringP("env", "e", "dev", `运行环境, 可选项目: dev or test prod`)
	pflag.StringP("config", "c", "app-dev.yaml", `配置文件路径`)
	pflag.Lookup("config").DefValue = "[./app-dev.yaml, ./config/app-dev.yaml]"
	pflag.Parse()

	// 命令行参数绑定
	if err := v.BindPFlags(pflag.CommandLine); err != nil {
		fmt.Println("命令行参数绑定失败")
		panic(err.Error())
	}

	// 环境变量参数绑定
	err := v.BindEnv("env", "RUN_ENV")
	if err != nil {
		fmt.Println("环境变量参数绑定失败")
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
		fmt.Printf("当前运行环境: %s\n", env)

	default:
		v.Set("env", "dev")
		v.SetConfigName("app-dev.yaml") // 默认dev环境
		fmt.Printf("未设置当前运行环境, 默认: %s\n", "dev")
	}

	// 设置viper, 加载配置文件
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	if err := v.ReadInConfig(); err != nil {
		fmt.Println("配置文件加载失败")
		panic(err.Error())
	}

	// 设置环境变量前缀为appName
	//v.SetEnvPrefix(v.GetString("appName"))

	fmt.Printf("使用配置文件%s\n", v.ConfigFileUsed())

}

func GetViper() *viper.Viper {
	return v
}
