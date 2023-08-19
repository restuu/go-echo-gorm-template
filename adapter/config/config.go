package config

import (
	"os"

	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	Port        int    `mapstructure:"PORT"`
	DatabaseURL string `mapstructure:"DB_URL"`
}

func NewConfig() (conf *Config, err error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	viper.AddConfigPath(cwd)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	conf = new(Config)

	err = viper.Unmarshal(conf)

	return
}
