package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Env        string `mapstructure:"ENVIRONMENT"`
	APIPort    string `mapstructure:"API_PORT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbUsername string `mapstructure:"DB_USERNAME"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName     string `mapstructure:"DB_NAME"`
	DbSSLMode  string `mapstructure:"DB_SSL_MODE"`
}

func Init() *Config {
	viper.AutomaticEnv()

	viper.BindEnv("ENVIRONMENT")
	viper.BindEnv("API_PORT")
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("DB_USERNAME")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DB_SSL_MODE")

	config := Config{}
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal config from env vars: %w", err))
	}

	return &config
}
