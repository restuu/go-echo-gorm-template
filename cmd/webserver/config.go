package main

import (
	"go-echo-gorm-tempate/pkg/config"

	"github.com/spf13/viper"
)

func newConfig() (conf *config.Config, err error) {
	conf = new(config.Config)

	err = viper.Unmarshal(conf)

	return
}
