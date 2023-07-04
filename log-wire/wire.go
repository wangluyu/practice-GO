//go:build wireinject
// +build wireinject

package log_wire

import "github.com/google/wire"

func InitLogger() (Logger, error) {
	wire.Build(ProvideSet)
	return Logger{}, nil
}
