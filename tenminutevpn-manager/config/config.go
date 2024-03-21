package config

import (
	"fmt"
	"os"

	"github.com/tenminutevpn/tenminutevpn-manager/resource"
	"github.com/tenminutevpn/tenminutevpn-manager/wireguard2"
	"gopkg.in/yaml.v3"
)

func ParseResources() ([]*resource.Resource, error) {
	file, err := os.Open("config/config-3.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var resources []*resource.Resource
	decoder := yaml.NewDecoder(file)

	for {
		var res resource.Resource
		if err := decoder.Decode(&res); err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		doc, err := yaml.Marshal(&res)
		if err != nil {
			return nil, err
		}

		switch res.Kind {
		case "wireguard/v1":
			var r wireguard2.Resource
			if err := yaml.Unmarshal(doc, &r); err != nil {
				return nil, fmt.Errorf("failed to parse wireguard resource: %w", err)
			}

			if err := r.Create(); err != nil {
				return nil, fmt.Errorf("failed to process wireguard resource: %w", err)
			}

			if err := r.Service().Enable(); err != nil {
				return nil, fmt.Errorf("failed to enable wireguard service: %w", err)
			}

			if err := r.Service().Start(); err != nil {
				return nil, fmt.Errorf("failed to start wireguard service: %w", err)
			}
		default:
			return nil, fmt.Errorf("unknown kind: %s", res.Kind)
		}

		resources = append(resources, &res)
	}

	return resources, nil
}
