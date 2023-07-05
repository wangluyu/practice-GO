package app

import (
	"github.com/google/wire"
	"practice/log"
)

type Application struct {
	Name   string
	Logger log.Loggers
}

func New(name string, logger log.Loggers) (*Application, error) {
	app := &Application{
		Name:   name,
		Logger: logger,
	}

	return app, nil
}

var ProvideSet = wire.NewSet(New)
