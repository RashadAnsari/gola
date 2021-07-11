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
		Commands []Command `mapstructure:"commands"`
	}

	// Command represents a shell command with its gola alias name.
	Command struct {
		Name string `mapstructure:"name"`
		Cmd  Cmd    `mapstructure:"cmd"`
	}

	// Cmd represents a shell command configuration.
	Cmd struct {
		Path string   `mapstructure:"path"`
		Args []string `mapstructure:"args"`
		Env  []string `mapstructure:"env"`
		Dir  string   `mapstructure:"dir"`
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
