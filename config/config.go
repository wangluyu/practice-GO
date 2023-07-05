package config

import (
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type ConfigPath string

func NewConfig(configPath ConfigPath) (*viper.Viper, error) {
	var (
		err error
		v   = viper.New()
	)
	v.AddConfigPath(".")
	v.SetConfigFile(string(configPath))

	if err := v.ReadInConfig(); err == nil {
		fmt.Printf("use config file -> %s\n", v.ConfigFileUsed())
	} else {
		return nil, err
	}

	return v, err
}

var ProvideSet = wire.NewSet(NewConfig)
