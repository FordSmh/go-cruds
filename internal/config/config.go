package config

import (
	"fmt"

	"github.com/vrischmann/envconfig"
)

type (
	Config struct {
		APIServerPort string `envconfig:"default=8009"`
		GinMode       string `envconfig:"default=debug"`
		LogLevel      string `envconfig:"LOG_LEVEL,default=warn"`
		Mongo         *Mongo
	}
)

func Init() (Config, error) {
	config := Config{}

	if err := envconfig.Init(&config); err != nil {
		return Config{}, fmt.Errorf("init config failed: %w", err)
	}

	return config, nil
}
