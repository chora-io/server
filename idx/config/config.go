package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config is the configuration.
type Config struct {
	// BackoffDuration is the duration between advancing processes.
	BackoffDuration time.Duration `mapstructure:"IDX_BACKOFF_DURATION"`

	// BackoffMaxRetries is the maximum number of retries before stopping a process.
	BackoffMaxRetries uint64 `mapstructure:"IDX_BACKOFF_MAX_RETRIES"`

	// ChainId is the chain id of the network (e.g. chora-testnet-1, regen-redwood-1).
	ChainId string `mapstructure:"IDX_CHAIN_ID"`

	// ChainRpc is the rpc endpoint for the network (e.g. testnet.chora.io:9090, redwood.chora.io:9090).
	ChainRpc string `mapstructure:"IDX_CHAIN_RPC"`

	// StartBlock is the starting block height from which the process will start when no record of the
	// process exists in the database. When a record does exist, start block is ignored.
	StartBlock int64 `mapstructure:"IDX_START_BLOCK"`

	// DatabaseUrl is the URL of the postgres database.
	DatabaseUrl string `mapstructure:"DATABASE_URL"`

	// ServerEnv is the environment the server is running in (i.e. local, staging, production).
	ServerEnv string `mapstructure:"SERVER_ENV"`
}

// LoadConfig loads the configuration.
func LoadConfig() Config {
	cfg := Config{}
	v := viper.New()
	v.SetDefault("DATABASE_URL", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	v.SetDefault("IDX_BACKOFF_DURATION", "5s")
	v.SetDefault("IDX_BACKOFF_MAX_RETRIES", 3)
	v.SetDefault("IDX_CHAIN_ID", "chora-local")
	v.SetDefault("IDX_CHAIN_RPC", "127.0.0.1:9090")
	v.SetDefault("IDX_START_BLOCK", 1)
	v.SetDefault("SERVER_ENV", "local")
	v.AutomaticEnv()
	if err := v.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	return cfg
}
