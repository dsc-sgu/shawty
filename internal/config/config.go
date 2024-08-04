package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var C appConfig

// Loads config from a YAML file.
func Load(configFile string) appConfig {
	content, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf(
			"Failed to load application config file %q. Error: %s\n",
			configFile,
			err,
		)
		panic("failed to load the config")
	}

	var cfg appConfig
	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		fmt.Printf(
			"Failed to unmarshal config file %s. Error %s\n",
			configFile,
			err,
		)
		panic("failed to load the config")
	}
	return cfg
}
