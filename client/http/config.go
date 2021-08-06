package http

import "github.com/kelseyhightower/envconfig"

type Config struct {
	MaxIdleConns        *int `default:"10"`
	MaxConnsPerHost     *int `default:"10"`
	MaxIdleConnsPerHost *int `default:"10"`
}

func New() *Config {
	var conf Config
	appName := "live-config"

	err := envconfig.Process(appName, &conf)

	if err != nil {
		panic(err)
	}

	return &conf
}
