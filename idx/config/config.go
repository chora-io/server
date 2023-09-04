package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config is the configuration.
type Config struct {
	// DatabaseUrl is the URL of the postgres database.
	DatabaseUrl string `mapstructure:"DATABASE_URL"`

	// IdxRunnerBackoff is the duration between advancing processes.
	IdxRunnerBackoff time.Duration `mapstructure:"IDX_RUNNER_BACKOFF"`

	// IdxRunnerMaxRetries is the maximum number of retries before stopping a process.
	IdxRunnerMaxRetries uint64 `mapstructure:"IDX_RUNNER_MAX_RETRIES"`

	// ServerEnv is the environment the server is running in (i.e. local, staging, production).
	ServerEnv string `mapstructure:"SERVER_ENV"`
}

// LoadConfig loads the configuration.
func LoadConfig() Config {
	cfg := Config{}
	v := viper.New()
	v.SetDefault("DATABASE_URL", "postgres://postgres:password@localhost:5432/server?sslmode=disable")
	v.SetDefault("IDX_RUNNER_BACKOFF", "1s")
	v.SetDefault("IDX_RUNNER_MAX_RETRIES", 10)
	v.SetDefault("SERVER_ENV", "local")
	v.AutomaticEnv()
	if err := v.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	return cfg
}
