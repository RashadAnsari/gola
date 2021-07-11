package config

import (
	"github.com/RashadAnsari/gola/pkg/config"
)

const (
	app       = "gola"
	cfgFile   = "gola.yaml"
	cfgPrefix = "gola"
)

type (
	// Config represents application configuration struct.
	Config struct {
	}
)

// Init initializes application configuration.
func Init() (*Config, error) {
	var cfg Config

	if err := config.Init(app, cfgFile, &cfg, defaultConfig, cfgPrefix); err != nil {
		return nil, err
	}

	return &cfg, nil
}
