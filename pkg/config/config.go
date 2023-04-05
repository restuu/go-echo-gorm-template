package config

// Config ...
type Config struct {
	Port        int    `mapstructure:"PORT"`
	DatabaseURL string `mapstructure:"DB_URL"`
}
