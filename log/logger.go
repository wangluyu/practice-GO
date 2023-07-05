package log

import (
	"encoding/json"
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Log interface {
	Debug()
	Info()
	Warn()
	Error()
	Panic()
	Fatal()
	With()
}

type Loggers map[string]*Logger

func (ls Loggers) Get(name string) (logger *Logger) {
	return ls[name]
}

type ZapLoggers map[string]*zap.SugaredLogger

type Options map[string]option

type Logger struct {
	name   string
	logger *zap.SugaredLogger
}

type option struct {
	LoggerName string `json:"logger_name"`
	LogPath    string `json:"log_path"`
	MaxSize    int    `json:"max_size"`
	MaxBackups int    `json:"max_backups"`
	MaxAge     int    `json:"max_age"`
	Level      string `json:"level"`
	Stdout     bool   `json:"stdout"`
}

func NewLoggerOptions(v *viper.Viper) (Options, error) {
	options := make(Options)
	logConfigs := v.Get("log")
	for _, item := range logConfigs.([]interface{}) {
		fmt.Println(item)
		j, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		o := new(option)
		err = json.Unmarshal(j, &o)
		if err != nil {
			return nil, err
		}
		options[o.LoggerName] = *o
	}
	return options, nil
}

func NewZapLoggers(os Options) (ZapLoggers, error) {
	zls := make(ZapLoggers)
	for name, option := range os {
		logger, err := zap.NewProduction()
		if err != nil {
			return nil, err
		}
		sugar := logger.Sugar()
		logger.Sync()
		fmt.Println(name)
		fmt.Println(option)
		zls[name] = sugar
	}
	return zls, nil
}

func NewLoggers(zapLoggers ZapLoggers) (Loggers, error) {
	loggers := make(Loggers)
	for name, zapLogger := range zapLoggers {
		loggers[name] = &Logger{name: name, logger: zapLogger}
	}
	return loggers, nil
}

func (l *Logger) Debug(msg string, keysAndValues ...interface{}) {
	l.logger.Debugf(msg, keysAndValues...)
}

func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

func (l *Logger) Warn(msg string, keysAndValues ...interface{}) {
	l.logger.Warnf(msg, keysAndValues...)
}

func (l *Logger) Error(msg string, keysAndValues ...interface{}) {
	l.logger.Errorf(msg, keysAndValues...)
}

func (l *Logger) With(args ...interface{}) *Logger {
	l.logger = l.logger.With(args...)
	return l
}

var ProvideSet = wire.NewSet(NewLoggers, NewZapLoggers, NewLoggerOptions)

// 注入依赖重点：不要在实现里 new
