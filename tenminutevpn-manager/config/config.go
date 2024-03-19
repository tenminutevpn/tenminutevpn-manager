package config

import (
	"fmt"
	"os"
)

type Config struct {
	Listeners []Listener `yaml:"listeners"`
}

type Listener struct {
	Name     string `yaml:"name"`
	Bind     Bind   `yaml:"bind"`
	Upstream string `yaml:"upstream"`
}

type Bind struct {
	Protocol string `yaml:"protocol"`
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
}

func parseConfig() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	fmt.Println(string(data))
	return nil, nil
}
