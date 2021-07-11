package config

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Init initializes a config struct using default, file, and environment variables.
func Init(app string, file string, cfg interface{}, defaultConfig string, prefix string) error {
	v := viper.New()
	v.SetConfigType("yaml")

	if err := v.ReadConfig(bytes.NewReader([]byte(defaultConfig))); err != nil {
		return err
	}

	v.SetConfigFile(file)
	v.SetEnvPrefix(prefix)
	v.AddConfigPath(fmt.Sprintf("/etc/%s/", app))
	v.AddConfigPath(fmt.Sprintf("$HOME/.%s", app))
	v.AddConfigPath(".")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	switch err := v.MergeInConfig(); err.(type) {
	case nil:
	case *os.PathError:
	default:
		return err
	}

	if err := v.UnmarshalExact(&cfg); err != nil {
		return err
	}

	return nil
}
