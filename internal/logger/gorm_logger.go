package logger

import (
	"context"
	"fmt"
	glogger "gorm.io/gorm/logger"
	"soul/global"
	"strings"
	"time"
)

type logrusAdapter struct {
	glogger.Config
	reportCaller bool
}

func (l logrusAdapter) LogMode(level glogger.LogLevel) glogger.Interface {
	return l
}

func (l logrusAdapter) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.reportCaller {
		log.WithContext(ctx).WithField("file", callerInfo(3)).Infof(msg, data...)
	} else {
		log.WithContext(ctx).Infof(msg, data...)
	}
}

func (l logrusAdapter) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.reportCaller {
		log.WithContext(ctx).WithField("file", callerInfo(3)).Warnf(msg, data...)
	} else {
		log.WithContext(ctx).Warnf(msg, data...)
	}
}

func (l logrusAdapter) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.reportCaller {
		log.WithContext(ctx).WithField("file", callerInfo(3)).Errorf(msg, data...)
	} else {
		log.WithContext(ctx).Errorf(msg, data...)
	}
}

func (l logrusAdapter) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	fmt.Println(ctx, begin, fc, err)
	//if l.LogLevel <= Silent {
	//	return
	//}
	//
	//elapsed := time.Since(begin)
	//switch {
	//case err != nil && l.Config.LogLevel >= glogger.Error && (!errors.Is(err, ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
	//	sql, rows := fc()
	//	if rows == -1 {
	//		l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
	//	} else {
	//		l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
	//	}
	//case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= Warn:
	//	sql, rows := fc()
	//	slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
	//	if rows == -1 {
	//		l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
	//	} else {
	//		l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
	//	}
	//case l.LogLevel == Info:
	//	sql, rows := fc()
	//	if rows == -1 {
	//		l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
	//	} else {
	//		l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
	//	}
	//}
}

func NewGormLogger() glogger.Interface {
	var logLevel glogger.LogLevel
	switch strings.ToLower(global.Config.GetString("database.logLevel")) {
	case "info":
		logLevel = glogger.Info
	case "warn":
		logLevel = glogger.Warn
	case "error":
		logLevel = glogger.Error
	case "silent":
		logLevel = glogger.Silent
	}
	gConfig := glogger.Config{
		SlowThreshold:             time.Second, // 慢 SQL 阈值
		LogLevel:                  logLevel,    // 日志级别
		IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
	}

	return logrusAdapter{
		Config:       gConfig,
		reportCaller: global.Config.GetBool("database.reportCaller"),
	}
}
