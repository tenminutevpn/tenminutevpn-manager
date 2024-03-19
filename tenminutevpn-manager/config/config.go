package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Metadata struct {
	Name string `yaml:"name"`
}

type Socket struct {
	Protocol string `yaml:"protocol"`
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
}

type Peer struct {
	Name      string `yaml:"name"`
	PublicKey string `yaml:"publicKey"`
}

type WireguardSpec struct {
	Socket     Socket `yaml:"socket"`
	PrivateKey string `yaml:"privateKey"`
	Peers      []Peer `yaml:"peers"`
}

type SquidSpec struct {
	Socket Socket `yaml:"socket"`
	Access string `yaml:"access"`
}

type Config struct {
	Kind     string      `yaml:"kind"`
	Metadata Metadata    `yaml:"metadata"`
	Spec     interface{} `yaml:"spec"`
}

func ParseConfig() ([]*Config, error) {
	file, err := os.Open("config/config-2.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var configs []*Config
	decoder := yaml.NewDecoder(file)

	for {
		var config Config
		err := decoder.Decode(&config)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		specData, err := yaml.Marshal(config.Spec)
		if err != nil {
			return nil, err
		}

		switch config.Kind {
		case "wireguard/v1":
			var wgSpec WireguardSpec
			err = yaml.Unmarshal(specData, &wgSpec)
			if err != nil {
				return nil, err
			}
			config.Spec = wgSpec
		case "squid/v1":
			var squidSpec SquidSpec
			err = yaml.Unmarshal(specData, &squidSpec)
			if err != nil {
				return nil, err
			}
			config.Spec = squidSpec
		default:
			return nil, fmt.Errorf("unknown kind: %s", config.Kind)
		}

		configs = append(configs, &config)
	}

	for _, config := range configs {
		fmt.Println(config)
	}
	return configs, nil
}
