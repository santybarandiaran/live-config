package redis

import (
	"github.com/kelseyhightower/envconfig"
	"time"
)

// Config represents a configuration object with values that can be passed via env variable
type Config struct {
	Address      *string        `default:":6379"`
	Password     *string        `default:""`
	GroupName    *string        `default:"gate-group"`
	ConsumerName *string        `default:"gate1" envconfig:"hostname"`
	StreamName   *string        `default:"ci-notifications"`
	ClaimMinIdle *time.Duration `default:"3m"`
	ClaimMax     *int           `default:"10"`
	Enabled      *bool          `default:"true"`
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
