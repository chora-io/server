package app

import "github.com/spf13/viper"

type Config struct {
	// DatabaseUrl is the URL of the postgres database.
	DatabaseUrl string `mapstructure:"DATABASE_URL"`

	// IdxPort is the port the application will run on.
	IdxPort uint64 `mapstructure:"IDX_PORT"`

	// ServerEnv is the environment the server is running in (i.e. local, staging, production).
	ServerEnv string `mapstructure:"SERVER_ENV"`
}

func LoadConfig() Config {
	cfg := Config{}
	v := viper.New()
	v.SetDefault("DATABASE_URL", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	v.SetDefault("IDX_PORT", 3001)
	v.SetDefault("SERVER_ENV", "local")
	v.AutomaticEnv()
	if err := v.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	return cfg
}
