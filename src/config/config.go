package config

import "github.com/spf13/viper"

var config *Config

type Config struct {
	FIRE_WATCH_PORT            string
	FIRE_WATCH_AUDIENCE        string
	FIRE_WATCH_ISSUER          string
	FIRE_WATCH_ACCESS_EXPIRED  int
	FIRE_WATCH_REFRESH_EXPIRED int
	FIRE_WATCH_JWT_SECRET      string
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
		FIRE_WATCH_PORT:            viper.GetString("FIRE_WATCH_PORT"),
		FIRE_WATCH_AUDIENCE:        viper.GetString("FIRE_WATCH_AUDIENCE"),
		FIRE_WATCH_ISSUER:          viper.GetString("FIRE_WATCH_ISSUER"),
		FIRE_WATCH_ACCESS_EXPIRED:  viper.GetInt("FIRE_WATCH_ACCESS_EXPIRED"),
		FIRE_WATCH_REFRESH_EXPIRED: viper.GetInt("FIRE_WATCH_REFRESH_EXPIRED"),
		FIRE_WATCH_JWT_SECRET:      viper.GetString("FIRE_WATCH_JWT_SECRET"),
	}
}

func GetCofig() *Config {
	return config
}
