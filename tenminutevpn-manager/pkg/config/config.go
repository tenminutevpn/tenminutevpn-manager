package config

import (
	"fmt"
	"os"

	"github.com/tenminutevpn/tenminutevpn-manager/pkg/resource"
	"github.com/tenminutevpn/tenminutevpn-manager/pkg/squid"
	"github.com/tenminutevpn/tenminutevpn-manager/pkg/wireguard"
	"gopkg.in/yaml.v3"
)

func ParseResources(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)

	for {
		var res resource.Resource
		if err := decoder.Decode(&res); err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}

		doc, err := yaml.Marshal(&res)
		if err != nil {
			return err
		}

		switch res.Kind {
		case "wireguard/v1":
			var r wireguard.Resource
			if err := yaml.Unmarshal(doc, &r); err != nil {
				return fmt.Errorf("failed to parse wireguard resource: %w", err)
			}

			if err := r.Create(); err != nil {
				return fmt.Errorf("failed to process wireguard resource: %w", err)
			}

			if err := r.Service().Enable(); err != nil {
				return fmt.Errorf("failed to enable wireguard service: %w", err)
			}

			if err := r.Service().Start(); err != nil {
				return fmt.Errorf("failed to start wireguard service: %w", err)
			}
		case "squid/v1":
			var r squid.Resource
			if err := yaml.Unmarshal(doc, &r); err != nil {
				return fmt.Errorf("failed to parse squid resource: %w", err)
			}

			if err := r.Create(); err != nil {
				return fmt.Errorf("failed to process squid resource: %w", err)
			}

			if err := r.Service().Enable(); err != nil {
				return fmt.Errorf("failed to enable squid service: %w", err)
			}

			if err := r.Service().Start(); err != nil {
				return fmt.Errorf("failed to start squid service: %w", err)
			}
		default:
			return fmt.Errorf("unknown kind: %s", res.Kind)
		}
	}

	return nil
}
