package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port     string `mapstructure:"PORT"`
	DBURL    string `mapstructure:"DATABASE_URL"`
	RedisURL string `mapstructure:"REDIS_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
