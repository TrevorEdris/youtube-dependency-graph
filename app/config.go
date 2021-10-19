package app

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Youtube YoutubeClientConfig
	Log     LogConfig
	Graph   GraphConfig
}

type YoutubeClientConfig struct {
	APIKey string `envconfig:"API_KEY" required:"true"`
}

type LogConfig struct {
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
	LogFmt   string `envconfig:"LOG_FMT" default:"logfmt"`
}

type GraphConfig struct {
	MaxDepth int `envconfig:"MAX_DEPTH" default:"3"`
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
