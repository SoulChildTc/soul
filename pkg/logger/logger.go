package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"soul/pkg/config"
)

var (
	log    *logrus.Entry
	v      *viper.Viper
	logger *logrus.Logger
)

func setOutPut(log *logrus.Logger) {
	var writer []io.Writer

	logPath := v.GetString("log.path")

	if v.GetBool("log.console") {
		writer = append(writer, os.Stdout)
	}

	if !v.GetBool("log.closeFileLog") {
		if v.GetBool("log.rotate.enable") {
			writer = append(writer, &lumberjack.Logger{
				Filename:   logPath,
				MaxSize:    v.GetInt("maxSize"),    // 单个文件最大大小, 单位M
				MaxBackups: v.GetInt("maxBackups"), // 最多保留多少个文件
				MaxAge:     v.GetInt("maxAge"),     // 每个最多保留多少天
				Compress:   v.GetBool("compress"),  // 启用压缩
				LocalTime:  v.GetBool("localTime"), // 默认使用UTC时间, 改为使用本地时间
			})
		} else {
			f, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
			if err != nil {
				fmt.Println("创建日志文件失败")
				panic(err.Error())
			}
			writer = append(writer, f)
		}
	}

	log.SetOutput(io.MultiWriter(writer...))
}

func InitLogger() {
	v = config.GetViper()
	logger = logrus.New()
	// 设置日志输出目标
	setOutPut(logger)

	// 设置日志格式
	if v.GetString("env") == "dev" {
		logger.SetFormatter(&logrus.TextFormatter{})
	} else {

		logger.SetFormatter(&logrus.JSONFormatter{})
	}

	level, err := logrus.ParseLevel(v.GetString("log.level"))
	if err != nil {
		panic("未知的日志级别,可选项为[TRACE, DEBUG, INFO, WARN, ERROR, FATAL, PANIC]")
	}
	logger.SetLevel(level)

	log = logger.WithFields(logrus.Fields{
		"service": v.GetString("appName"),
	})
}

func GetLogger() *logrus.Logger {
	return logger
}

func GetEntry() *logrus.Entry {
	return log
}

func Trace(format string, args ...interface{}) {
	log.Tracef(format, args...)
}

func Debug(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Info(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warn(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Error(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Fatal(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Panic(format string, args ...interface{}) {
	log.Panicf(format, args...)
}
