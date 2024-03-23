package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tenminutevpn/tenminutevpn-manager/pkg/provider/dnsmasq"
	"github.com/tenminutevpn/tenminutevpn-manager/pkg/provider/squid"
	"github.com/tenminutevpn/tenminutevpn-manager/pkg/provider/wireguard"
	"github.com/tenminutevpn/tenminutevpn-manager/pkg/resource"
	"gopkg.in/yaml.v3"
)

func Parse(path string) error {
	// check if path exists
	file, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("path not found: %s", path)
	}

	if file.IsDir() {
		return ParseResourcesDirectory(path)
	} else {
		return ParseResources(path)
	}

}

func ParseResourcesDirectory(directory string) error {
	files, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) != ".yaml" {
			continue
		}

		filename := filepath.Join(directory, file.Name())
		if err := ParseResources(filename); err != nil {
			return err
		}
	}
	return nil
}

func ParseResources(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("Parsing file:", filename)

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

		fmt.Println("Resource:", res.Kind, res.Metadata.Name)

		switch res.Kind {
		case "dnsmasq/v1":
			var r dnsmasq.Resource
			if err := yaml.Unmarshal(doc, &r); err != nil {
				return fmt.Errorf("failed to parse dnsmasq resource: %w", err)
			}

			if err := r.Create(); err != nil {
				return fmt.Errorf("failed to process dnsmasq resource: %w", err)
			}

			if err := r.Service().Enable(); err != nil {
				return fmt.Errorf("failed to enable dnsmasq service: %w", err)
			}

			if err := r.Service().Start(); err != nil {
				return fmt.Errorf("failed to start dnsmasq service: %w", err)
			}

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
