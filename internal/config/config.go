package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	ServerPort string

	AgifyURL       string
	GenderizeURL   string
	NationalizeURL string

	HTTPTimeout         time.Duration
	HTTPMaxIdleConns    int
	HTTPIdleConnTimeout time.Duration

	Environment string
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	viper.AutomaticEnv()

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "user")
	viper.SetDefault("DB_PASSWORD", "password")
	viper.SetDefault("DB_NAME", "people_db")

	viper.SetDefault("SERVER_PORT", "8081")

	viper.SetDefault("AGIFY_URL", "https://api.agify.io")
	viper.SetDefault("GENDERIZE_URL", "https://api.genderize.io")
	viper.SetDefault("NATIONALIZE_URL", "https://api.nationalize.io")

	viper.SetDefault("HTTP_TIMEOUT", "10s")
	viper.SetDefault("HTTP_MAX_IDLE_CONNS", 100)
	viper.SetDefault("HTTP_IDLE_CONN_TIMEOUT", "90s")

	viper.SetDefault("ENV", "development")

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	httpTimeout, err := time.ParseDuration(viper.GetString("HTTP_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("invalid HTTP_TIMEOUT: %w", err)
	}
	cfg.HTTPTimeout = httpTimeout

	httpIdleTimeout, err := time.ParseDuration(viper.GetString("HTTP_IDLE_CONN_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("invalid HTTP_IDLE_CONN_TIMEOUT: %w", err)
	}
	cfg.HTTPIdleConnTimeout = httpIdleTimeout

	return &cfg, nil
}
