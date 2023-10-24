package auth

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Configuration struct {
	DBUser                 string        `mapstructure:"DB_USER"`
	DBPassword             string        `mapstructure:"DB_PASSWORD"`
	DBHost                 string        `mapstructure:"DB_HOST"`
	DBPort                 int32         `mapstructure:"DB_PORT"`
	DBName                 string        `mapstructure:"DB_NAME"`
	HttpPort               string        `mapstructure:"HTTP_PORT"`
	GinMode                string        `mapstructure:"GIN_MODE"`
	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`
}

func NewConfig() *Configuration {
	var config Configuration
	viper.SetConfigFile("config/auth/app.env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return &config
}
