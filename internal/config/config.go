package config

import (
	"fmt"
	"github.com/SoulChildTc/soul/global"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	// 初始化viper
	v := viper.New()

	// 命令行参数获取
	pflag.StringP("config", "c", "app.yaml", `配置文件路径`)
	pflag.BoolP("migrate", "m", false, `迁移数据库`)
	pflag.Lookup("config").DefValue = "[./app.yaml, ./config/app.yaml]"
	pflag.Parse()

	// 命令行参数绑定
	if err := v.BindPFlags(pflag.CommandLine); err != nil {
		fmt.Println("[Init] 命令行参数绑定失败")
		panic(err.Error())
	}

	// 设置默认参数
	setDefaultParams(v)

	// 读取配置文件
	configFilePath := v.GetString("config")
	if configFilePath != "" {
		v.SetConfigName(configFilePath)
	} else {
		v.SetConfigName("app.yaml")
	}

	// 设置viper, 加载配置文件
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("[Warning] 配置文件读取失败: %s, 使用默认配置\n", err.Error())
	} else {
		fmt.Printf("[Init] 使用配置文件: %s\n", v.ConfigFileUsed())
	}

	err := v.Unmarshal(&global.Config)
	if err != nil {
		panic("加载配置失败" + err.Error())
	}

	return v
}
