package app

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Youtube YoutubeClientConfig
}

type YoutubeClientConfig struct {
	APIKey string `envconfig:"API_KEY" required:"true"`
}

func ParseConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		return Config{}, err
	}

	var cfg Config
	err = envconfig.Process("", &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
