package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	ServerPort string `mapstructure:"SERVER_PORT"`

	AgifyURL       string `mapstructure:"AGIFY_URL"`
	GenderizeURL   string `mapstructure:"GENDERIZE_URL"`
	NationalizeURL string `mapstructure:"NATIONALIZE_URL"`

	HTTPTimeout         time.Duration `mapstructure:"HTTP_TIMEOUT"`
	HTTPMaxIdleConns    int           `mapstructure:"HTTP_MAX_IDLE_CONNS"`
	HTTPIdleConnTimeout time.Duration `mapstructure:"HTTP_IDLE_CONN_TIMEOUT"`

	Environment string `mapstructure:"ENV"`
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "postgres")
	viper.SetDefault("DB_NAME", "people_db")

	viper.SetDefault("SERVER_PORT", "8080")

	viper.SetDefault("AGIFY_URL", "https://api.agify.io")
	viper.SetDefault("GENDERIZE_URL", "https://api.genderize.io")
	viper.SetDefault("NATIONALIZE_URL", "https://api.nationalize.io")

	viper.SetDefault("HTTP_TIMEOUT", "10s")
	viper.SetDefault("HTTP_MAX_IDLE_CONNS", 100)
	viper.SetDefault("HTTP_IDLE_CONN_TIMEOUT", "90s")

	viper.SetDefault("ENV", "development")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading .env file: %w", err)
		}
	}

	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}

	var err error
	cfg.HTTPTimeout, err = time.ParseDuration(viper.GetString("HTTP_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("invalid HTTP_TIMEOUT: %w", err)
	}

	cfg.HTTPIdleConnTimeout, err = time.ParseDuration(viper.GetString("HTTP_IDLE_CONN_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf("invalid HTTP_IDLE_CONN_TIMEOUT: %w", err)
	}

	return &cfg, nil
}
