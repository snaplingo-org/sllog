package sllog

import (
	"os"

	"github.com/sirupsen/logrus"
)

type gormlogger struct{}

type sllogger struct{}

var (
	sllog      = &sllogger{}
	gormLogger = &gormlogger{}
	logger     = logrus.New()
)

func (*gormlogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		logger.WithFields(logrus.Fields{"module": "gorm", "type": "sql"}).Print(v[:])
	}
	if v[0] == "log" {
		logger.WithFields(logrus.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}

// GetLogger get sllog logger
func GetLogger() *sllogger {
	return sllog
}

// GetGormLogger get gorm logger
func GetGormLogger() *gormlogger {
	return gormLogger
}

// GetLogrusLogger get logrus logger
func GetLogrusLogger() *logrus.Logger {
	return logger
}

/*
example:
FileLogConfig{
        Path:     "./log"
		Filename: "./log/app.log",
		MaxLines: 10000000000000,
		Maxsize:  10000000000000,
		Daily:    true,
		MaxDays:  3,
		Rotate:   true,
		Level:    LevelDebug,
}
*/
func Init(config FileLogConfig) {
	os.MkdirAll(config.Path, 0777)
	hook, err := newFileHook(config)
	if err != nil {
		logger.Error(err)
	}
	logger.Level = logrus.DebugLevel
	logger.AddHook(hook)
}

// Debug logs a message with debug log level.
func (sllog *sllogger) Debug(msg interface{}) {
	logger.Debug(msg)
}

// DebugWithF logs a message with Debug log level.
func (sllog *sllogger) DebugWithF(json *map[string]interface{}, msg interface{}) {
	if json == nil {
		logger.Debug(msg)
	} else {
		logger.WithFields(*json).Debug(msg)
	}
}

// Debugf logs a formatted message with debug log level.
func (sllog *sllogger) Debugf(msg string, args ...interface{}) {
	logger.Debugf(msg, args...)
}

// DebugfWithF logs a message with Debug log level.
func (sllog *sllogger) DebugfWithF(json *map[string]interface{}, msg string, args ...interface{}) {
	if json == nil {
		logger.Debugf(msg, args...)
	} else {
		logger.WithFields(*json).Debugf(msg, args...)
	}
}

// Info logs a message with info log level.
func (sllog *sllogger) Info(msg interface{}) {
	logger.Info(msg)
}

// InfoWithF logs a message with info log level.
func (sllog *sllogger) InfoWithF(json *map[string]interface{}, msg interface{}) {
	if json == nil {
		logger.Info(msg)
	} else {
		logger.WithFields(*json).Info(msg)
	}
}

// Infof logs a formatted message with info log level.
func (sllog *sllogger) Infof(msg string, args ...interface{}) {
	logger.Infof(msg, args...)
}

// InfofWithF logs a message with info log level.
func (sllog *sllogger) InfofWithF(json *map[string]interface{}, msg string, args ...interface{}) {
	if json == nil {
		logger.Infof(msg, args...)
	} else {
		logger.WithFields(*json).Infof(msg, args...)
	}
}

// Warn logs a message with warn log level.
func (sllog *sllogger) Warn(msg interface{}) {
	logger.Warn(msg)
}

// WarnWithF logs a message with info log level.
func (sllog *sllogger) WarnWithF(json *map[string]interface{}, msg interface{}) {
	if json == nil {
		logger.Warn(msg)
	} else {
		logger.WithFields(*json).Warn(msg)
	}
}

// Warnf logs a formatted message with warn log level.
func (sllog *sllogger) Warnf(msg string, args ...interface{}) {
	logger.Warnf(msg, args...)
}

// WarnfWithF logs a message with warn log level.
func (sllog *sllogger) WarnfWithF(json *map[string]interface{}, msg string, args ...interface{}) {
	if json == nil {
		logger.Warnf(msg, args...)
	} else {
		logger.WithFields(*json).Warnf(msg, args...)
	}
}

// Error logs a message with error log level.
func (sllog *sllogger) Error(msg interface{}) {
	logger.Error(msg)
}

// ErrorWithF logs a message with error log level.
func (sllog *sllogger) ErrorWithF(json *map[string]interface{}, msg interface{}) {
	if json == nil {
		logger.Error(msg)
	} else {
		logger.WithFields(*json).Error(msg)
	}
}

// Errorf logs a formatted message with error log level.
func (sllog *sllogger) Errorf(msg string, args ...interface{}) {
	logger.Errorf(msg, args...)
}

// ErrorfWithF logs a message with error log level.
func (sllog *sllogger) ErrorfWithF(json *map[string]interface{}, msg string, args ...interface{}) {
	if json == nil {
		logger.Errorf(msg, args...)
	} else {
		logger.WithFields(*json).Errorf(msg, args...)
	}
}

// Fatal logs a message with fatal log level.
func (sllog *sllogger) Fatal(msg interface{}) {
	logger.Fatal(msg)
}

// FatalWithF logs a message with fatal log level.
func (sllog *sllogger) FatalWithF(json *map[string]interface{}, msg interface{}) {
	if json == nil {
		logger.Fatal(msg)
	} else {
		logger.WithFields(*json).Fatal(msg)
	}
}

// Fatalf logs a formatted message with fatal log level.
func (sllog *sllogger) Fatalf(msg string, args ...interface{}) {
	logger.Fatalf(msg, args...)
}

// FatalfWithF logs a message with fatal log level.
func (sllog *sllogger) FatalfWithF(json *map[string]interface{}, msg string, args ...interface{}) {
	if json == nil {
		logger.Fatalf(msg, args...)
	} else {
		logger.WithFields(*json).Fatalf(msg, args...)
	}
}
