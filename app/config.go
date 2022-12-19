package app

import (
	"github.com/spf13/viper"
)

type Config struct {
	// AppEnv is the environment the server is running in (i.e. local, staging, production).
	AppEnv string `mapstructure:"APP_ENV"`

	// AppPort is the port the server will run on.
	AppPort uint64 `mapstructure:"APP_PORT"`

	// AppDatabaseUrl is the URL of the postgres database.
	AppDatabaseUrl string `mapstructure:"APP_DATABASE_URL"`

	// AppAllowedOrigins are the allowed origins for cross-origin requests.
	AppAllowedOrigins string `mapstructure:"APP_ALLOWED_ORIGINS"`
}

func LoadConfig() Config {
	cfg := Config{}
	v := viper.New()
	v.SetDefault("APP_ENV", "local")
	v.SetDefault("APP_PORT", 3000)
	v.SetDefault("APP_DATABASE_URL", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	v.SetDefault("APP_ALLOWED_ORIGINS", "*")
	v.AutomaticEnv()
	if err := v.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	return cfg
}
