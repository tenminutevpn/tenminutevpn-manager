package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Listeners []Listener `yaml:"listeners"`
	Upstreams []Upstream `yaml:"upstreams"`
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

type Upstream struct {
	Name string `yaml:"name"`
	Kind string `yaml:"kind"`
	Spec any    `yaml:"spec"`
}

type WireguardSpec struct {
	PrivateKey string `yaml:"privateKey"`
}

func ParseConfig() (*Config, error) {
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	for _, upstream := range config.Upstreams {
		switch upstream.Kind {
		case "wireguard":
			specBytes, err := yaml.Marshal(upstream.Spec) // Re-marshal the spec part
			if err != nil {
				return nil, err
			}

			var wgSpec WireguardSpec
			err = yaml.Unmarshal(specBytes, &wgSpec)
			if err != nil {
				return nil, err
			}

			fmt.Println(wgSpec.PrivateKey)
			// config.Upstreams[i].Spec = wgSpec
			// fmt.Println(wgSpec)
			// default:
			// 	return nil, fmt.Errorf("unknown upstream kind: %s", upstream.Kind)
		}
	}

	// fmt.Println(config)
	return nil, nil
}
