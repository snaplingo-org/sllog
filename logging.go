package sllog

import (
	"os"

	"github.com/sirupsen/logrus"
)

type GormLogger struct{}

type Sllog struct{}

var (
	sllog      = &Sllog{}
	gormLogger = &GormLogger{}
	logger     = logrus.New()
)

func (*GormLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		logger.WithFields(logrus.Fields{"module": "gorm", "type": "sql"}).Print(v[:])
	}
	if v[0] == "log" {
		logger.WithFields(logrus.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}

// GetLogger get sllog logger
func GetLogger() *Sllog {
	return sllog
}

// GetGormLogger get gorm logger
func GetGormLogger() *GormLogger {
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
	hook, err := NewFileHook(config)
	if err != nil {
		logger.Error(err)
	}
	logger.Level = logrus.DebugLevel
	logger.AddHook(hook)
}

// Debug logs a message with debug log level.
func (sllog *Sllog) Debug(msg interface{}) {
	logger.Debug(msg)
}

// DebugWithF logs a message with Debug log level.
func (sllog *Sllog) DebugWithF(json *map[string]interface{}, msg interface{}) {
	if json == nil {
		logger.Debug(msg)
	} else {
		logger.WithFields(*json).Debug(msg)
	}
}

// Debugf logs a formatted message with debug log level.
func (sllog *Sllog) Debugf(msg string, args ...interface{}) {
	logger.Debugf(msg, args...)
}

// DebugfWithF logs a message with Debug log level.
func (sllog *Sllog) DebugfWithF(json *map[string]interface{}, msg string, args ...interface{}) {
	if json == nil {
		logger.Debugf(msg, args...)
	} else {
		logger.WithFields(*json).Debugf(msg, args...)
	}
}

// Info logs a message with info log level.
func (sllog *Sllog) Info(msg interface{}) {
	logger.Info(msg)
}

// InfoWithF logs a message with info log level.
func (sllog *Sllog) InfoWithF(json *map[string]interface{}, msg interface{}) {
	if json == nil {
		logger.Info(msg)
	} else {
		logger.WithFields(*json).Info(msg)
	}
}

// Infof logs a formatted message with info log level.
func (sllog *Sllog) Infof(msg string, args ...interface{}) {
	logger.Infof(msg, args...)
}

// InfofWithF logs a message with info log level.
func (sllog *Sllog) InfofWithF(json *map[string]interface{}, msg string, args ...interface{}) {
	if json == nil {
		logger.Infof(msg, args...)
	} else {
		logger.WithFields(*json).Infof(msg, args...)
	}
}

// Warn logs a message with warn log level.
func (sllog *Sllog) Warn(msg interface{}) {
	logger.Warn(msg)
}

// WarnWithF logs a message with info log level.
func (sllog *Sllog) WarnWithF(json *map[string]interface{}, msg interface{}) {
	if json == nil {
		logger.Warn(msg)
	} else {
		logger.WithFields(*json).Warn(msg)
	}
}

// Warnf logs a formatted message with warn log level.
func (sllog *Sllog) Warnf(msg string, args ...interface{}) {
	logger.Warnf(msg, args...)
}

// WarnfWithF logs a message with warn log level.
func (sllog *Sllog) WarnfWithF(json *map[string]interface{}, msg string, args ...interface{}) {
	if json == nil {
		logger.Warnf(msg, args...)
	} else {
		logger.WithFields(*json).Warnf(msg, args...)
	}
}

// Error logs a message with error log level.
func (sllog *Sllog) Error(msg interface{}) {
	logger.Error(msg)
}

// ErrorWithF logs a message with error log level.
func (sllog *Sllog) ErrorWithF(json *map[string]interface{}, msg interface{}) {
	if json == nil {
		logger.Error(msg)
	} else {
		logger.WithFields(*json).Error(msg)
	}
}

// Errorf logs a formatted message with error log level.
func (sllog *Sllog) Errorf(msg string, args ...interface{}) {
	logger.Errorf(msg, args...)
}

// ErrorfWithF logs a message with error log level.
func (sllog *Sllog) ErrorfWithF(json *map[string]interface{}, msg string, args ...interface{}) {
	if json == nil {
		logger.Errorf(msg, args...)
	} else {
		logger.WithFields(*json).Errorf(msg, args...)
	}
}

// Fatal logs a message with fatal log level.
func (sllog *Sllog) Fatal(msg interface{}) {
	logger.Fatal(msg)
}

// FatalWithF logs a message with fatal log level.
func (sllog *Sllog) FatalWithF(json *map[string]interface{}, msg interface{}) {
	if json == nil {
		logger.Fatal(msg)
	} else {
		logger.WithFields(*json).Fatal(msg)
	}
}

// Fatalf logs a formatted message with fatal log level.
func (sllog *Sllog) Fatalf(msg string, args ...interface{}) {
	logger.Fatalf(msg, args...)
}

// FatalfWithF logs a message with fatal log level.
func (sllog *Sllog) FatalfWithF(json *map[string]interface{}, msg string, args ...interface{}) {
	if json == nil {
		logger.Fatalf(msg, args...)
	} else {
		logger.WithFields(*json).Fatalf(msg, args...)
	}
}
