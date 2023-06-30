package app

import "github.com/spf13/viper"

type Config struct {
	// ApiAllowedOrigins are the allowed origins for cross-origin requests.
	ApiAllowedOrigins string `mapstructure:"API_ALLOWED_ORIGINS"`

	// ApiPort is the port the application will run on.
	ApiPort uint64 `mapstructure:"API_PORT"`

	// DatabaseUrl is the URL of the postgres database.
	DatabaseUrl string `mapstructure:"DATABASE_URL"`

	// ServerEnv is the environment the server is running in (i.e. local, staging, production).
	ServerEnv string `mapstructure:"SERVER_ENV"`
}

func LoadConfig() Config {
	cfg := Config{}
	v := viper.New()
	v.SetDefault("API_ALLOWED_ORIGINS", "*")
	v.SetDefault("API_PORT", 3000)
	v.SetDefault("DATABASE_URL", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	v.SetDefault("SERVER_ENV", "local")
	v.AutomaticEnv()
	if err := v.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	return cfg
}
