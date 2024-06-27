package config

import "github.com/spf13/viper"

var config *Config

type Config struct {
	POSTGRES_PASSWORD string
	POSTGRES_USER     string
	POSTGRES_DB       string
	POSTGRES_HOST     string
	POSTGRES_PORT     string

	FIRE_WATCH_API_PORT string

	JWT_AUDIENCE        string
	JWT_ISSUER          string
	JWT_ACCESS_EXPIRED  int
	JWT_REFRESH_EXPIRED int
	JWT_SECRET          string

	SMTP_HOST_EMAIL    string
	SMTP_HOST          string
	SMTP_PORT          string
	SMTP_HOST_USER     string
	SMTP_HOST_PASSWORD string

	BLOB_ACCESS_KEY  string
	BLOB_PROJECT_KEY string
	BLOB_PUBLIC_URL  string
	BLOB_S3_URL      string
	BLOB_REGION      string

	REDIS_USER   string
	REDIS_PASSWD string
	REDIS_HOST   string
	REDIS_PORT   string
	REDIS_DB     int
}

func LoadEnv(path string) {
	viper.SetConfigName(path)
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		viper.Reset()
		viper.AutomaticEnv()
	}

	config = &Config{
		POSTGRES_PASSWORD: viper.GetString("POSTGRES_PASSWORD"),
		POSTGRES_USER:     viper.GetString("POSTGRES_USER"),
		POSTGRES_DB:       viper.GetString("POSTGRES_DB"),
		POSTGRES_HOST:     viper.GetString("POSTGRES_HOST"),
		POSTGRES_PORT:     viper.GetString("POSTGRES_PORT"),

		FIRE_WATCH_API_PORT: viper.GetString("FIRE_WATCH_API_PORT"),

		JWT_AUDIENCE:        viper.GetString("JWT_AUDIENCE"),
		JWT_ISSUER:          viper.GetString("JWT_ISSUER"),
		JWT_ACCESS_EXPIRED:  viper.GetInt("JWT_ACCESS_EXPIRED"),
		JWT_REFRESH_EXPIRED: viper.GetInt("JWT_REFRESH_EXPIRED"),
		JWT_SECRET:          viper.GetString("JWT_SECRET"),

		SMTP_HOST_EMAIL:    viper.GetString("SMTP_HOST_EMAIL"),
		SMTP_HOST:          viper.GetString("SMTP_HOST"),
		SMTP_PORT:          viper.GetString("SMTP_PORT"),
		SMTP_HOST_USER:     viper.GetString("SMTP_HOST_USER"),
		SMTP_HOST_PASSWORD: viper.GetString("SMTP_HOST_PASSWORD"),

		BLOB_ACCESS_KEY:  viper.GetString("BLOB_ACCESS_KEY"),
		BLOB_PROJECT_KEY: viper.GetString("BLOB_PROJECT_KEY"),
		BLOB_PUBLIC_URL:  viper.GetString("BLOB_PUBLIC_URL"),
		BLOB_S3_URL:      viper.GetString("BLOB_S3_URL"),
		BLOB_REGION:      viper.GetString("BLOB_REGION"),

		REDIS_USER:   viper.GetString("REDIS_USER"),
		REDIS_PASSWD: viper.GetString("REDIS_PASSWD"),
		REDIS_HOST:   viper.GetString("REDIS_HOST"),
		REDIS_PORT:   viper.GetString("REDIS_PORT"),
		REDIS_DB:     viper.GetInt("REDIS_DB"),
	}
}

func GetCofig() *Config {
	return config
}
