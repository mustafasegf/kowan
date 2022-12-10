package util

import (
	"errors"
	"os"
)

type Config struct {
	Port     string `mapstructure:"PORT"`
	DBURL    string `mapstructure:"DATABASE_URL"`
	RedisURL string `mapstructure:"REDIS_URL"`
}

func LoadConfig() (config Config, err error) {
	config = Config{
		Port:     os.Getenv("PORT"),
		DBURL:    os.Getenv("DATABASE_URL"),
		RedisURL: os.Getenv("REDIS_URL"),
	}
	if config.Port == "" {
		return config, errors.New("PORT is not set")
	}
	if config.DBURL == "" {
		return config, errors.New("DATABASE_URL is not set")
	}
	if config.RedisURL == "" {
		return config, errors.New("REDIS_URL is not set")
	}

	return
}
