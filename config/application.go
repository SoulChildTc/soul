package config

type Configuration struct {
	AppName  string   `yaml:"appName" mapstructure:"appName"`
	Listen   string   `yaml:"listen" mapstructure:"listen"`
	Port     int      `yaml:"port" mapstructure:"port"`
	Mode     string   `yaml:"mode" mapstructure:"mode"`
	Log      Log      `yaml:"log" mapstructure:"log"`
	Config   string   `yaml:"config" mapstructure:"config"`
	Database Database `yaml:"database" mapstructure:"database"`
	Jwt      Jwt      `yaml:"jwt" mapstructure:"jwt"`
}
