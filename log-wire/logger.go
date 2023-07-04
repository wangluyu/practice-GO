package log_wire

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

type Log interface {
}

type Logger struct {
	name   string
	logger *zap.SugaredLogger
}

type LogOption struct {
	name string
}

func newLoggerOption() (LogOption, error) {
	return LogOption{name: "test"}, nil
}

func newZapLogger(lo LogOption) (*zap.SugaredLogger, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	return sugar, nil
}

func newLogger(logger *zap.SugaredLogger) (Logger, error) {
	return Logger{name: "Test", logger: logger}, nil
}

func (l Logger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

var ProvideSet = wire.NewSet(newLogger, newZapLogger, newLoggerOption)

// 注入依赖重点：不要在非工厂的函数体内 new
