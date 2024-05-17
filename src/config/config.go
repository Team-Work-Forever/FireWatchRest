package config

import "github.com/spf13/viper"

var config *Config

type Config struct {
	FIRE_WATCH_PORT string
}

func LoadEnv(path string) {
	viper.SetConfigName(path)
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		viper.Reset()
	}

	config = &Config{
		FIRE_WATCH_PORT: viper.GetString("FIRE_WATCH_PORT"),
	}
}

func GetCofig() *Config {
	return config
}
