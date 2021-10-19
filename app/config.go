package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	enforcedMaximumDepth = 10
)

var (
	errEmptyAPIKey   = errors.New("missing required environment variable API_KEY")
	errEmptyLogFmt   = errors.New("missing required environment variable LOG_FMT")
	errEmptyLogLevel = errors.New("missing required environment variable LOG_LEVEL")
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

	err = cfg.Validate()
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (cfg Config) Validate() error {
	err := cfg.Graph.Validate()
	if err != nil {
		return err
	}

	err = cfg.Log.Validate()
	if err != nil {
		return err
	}

	err = cfg.Youtube.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (yCfg YoutubeClientConfig) Validate() error {
	if yCfg.APIKey == "" {
		return errEmptyAPIKey
	}
	return nil
}

func (lCfg LogConfig) Validate() error {
	if lCfg.LogFmt == "" {
		return errEmptyLogFmt
	}
	if lCfg.LogLevel == "" {
		return errEmptyLogLevel
	}
	return nil
}

func (gCfg GraphConfig) Validate() error {
	if gCfg.MaxDepth > enforcedMaximumDepth {
		return fmt.Errorf("provided MAX_DEPTH (%d) too high; Must be lower than %d", gCfg.MaxDepth, enforcedMaximumDepth)
	}
	return nil
}
