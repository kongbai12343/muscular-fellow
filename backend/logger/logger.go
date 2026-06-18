package logger

import "go.uber.org/zap"

var logger *zap.SugaredLogger

func Init() error {
	// 1. 创建一个生产环境配置的 zap.Logger
	log, err := zap.NewProduction()
	if err != nil {
		return err
	}

	// 2. 通过 .Sugar() 方法获取 SugaredLogger
	logger = log.Sugar()
	return nil
}

// 使用 defer 确保日志缓冲区的内容在程序退出前被写入
func Sync() error {
	if logger == nil {
		return nil
	}
	return logger.Sync()
}

// 直接输出日志
func Debug(args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Debug(args...)
}
func Info(args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Info(args...)
}
func Warn(args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Warn(args...)
}
func Error(args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Error(args...)
}

// 格式化输出日志
func Infof(template string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Infof(template, args...)
}

func Errorf(template string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Errorf(template, args...)
}

func Debugf(template string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Debugf(template, args...)
}

func Warnf(template string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Warnf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Fatalf(template, args...)
}

// 带键值对的输出日志
func Infow(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}
	logger.Infow(msg, keysAndValues...)
}
