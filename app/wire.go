//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"practice/config"
	"practice/log"
)

var AppProvideSet = wire.NewSet(
	log.ProvideSet,
	ProvideSet,
	config.ProvideSet,
)

func InitApp(cp config.ConfigPath, name string) (*Application, error) {
	wire.Build(AppProvideSet)
	return &Application{}, nil
}
