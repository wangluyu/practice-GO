//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
)

func InitEvent(phrase string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
