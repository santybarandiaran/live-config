package db

import (
	"github.com/kelseyhightower/envconfig"
)

// Config represents a configuration object with values that can be passed via env variable
type Config struct {
	DBHost            *string        `default:"localhost"`
	DBPort            *int           `default:"5432"`
	DBUser            *string        `default:"config"`
	DBPassword        *string        `default:"config"`
	DBName            *string        `default:"live_config"`
	DBSSLMode         *string        `default:"disable"`
	DBMaxIdleConns    *int           `default:"10"  envconfig:"db_max_idle_conns"`
	DBMaxOpenConns    *int           `default:"100" envconfig:"db_max_open_conns"`
}

// New gets a configuration object
func New() *Config {
	var conf Config
	appName := "live-config"

	err := envconfig.Process(appName, &conf)

	if err != nil {
		panic(err)
	}

	return &conf
}
