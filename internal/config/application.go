package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type ApplicationConfig struct {
	Queue QueueConfig
}

//ReadConfig reads configuration from the environment and populate the structure with it
func ReadConfig() (*ApplicationConfig, error) {
	var conf ApplicationConfig
	if err := envconfig.Process("", &conf); err != nil {
		return nil, err
	}
	fmt.Printf("conf: %+v\n", conf)
	return &conf, nil
}
