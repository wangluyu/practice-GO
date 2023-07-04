// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package log_wire

// Injectors from wire.go:

func InitLogger() (Logger, error) {
	logOption, err := newLoggerOption()
	if err != nil {
		return Logger{}, err
	}
	sugaredLogger, err := newZapLogger(logOption)
	if err != nil {
		return Logger{}, err
	}
	logger, err := newLogger(sugaredLogger)
	if err != nil {
		return Logger{}, err
	}
	return logger, nil
}