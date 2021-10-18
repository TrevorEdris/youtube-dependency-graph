package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type YoutubeClientConfig struct {
	APIKey string `envconfig:"API_KEY" required:"true"`
}

func parseConfig() (YoutubeClientConfig, error) {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		return YoutubeClientConfig{}, err
	}

	var cfg YoutubeClientConfig
	err = envconfig.Process("", &cfg)
	if err != nil {
		return YoutubeClientConfig{}, err
	}

	return cfg, nil
}
