package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config is the configuration.
type Config struct {
	// RunnerBackoffDuration is the duration between advancing processes.
	RunnerBackoffDuration time.Duration `mapstructure:"IDX_RUNNER_BACKOFF_DURATION"`

	// RunnerBackoffMaxRetries is the maximum number of retries before stopping a process.
	RunnerBackoffMaxRetries uint64 `mapstructure:"IDX_RUNNER_BACKOFF_MAX_RETRIES"`
}

// LoadConfig loads the configuration.
func LoadConfig() Config {
	cfg := Config{}
	v := viper.New()
	v.SetDefault("IDX_RUNNER_BACKOFF_DURATION", "500ms")
	v.SetDefault("IDX_RUNNER_BACKOFF_MAX_RETRIES", 3)
	v.AutomaticEnv()
	if err := v.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	return cfg
}
