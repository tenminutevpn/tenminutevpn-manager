package config

import (
	"os"

	"github.com/tenminutevpn/tenminutevpn-manager/resource"
	"gopkg.in/yaml.v3"
)

func ParseResources() ([]*resource.Resource, error) {
	file, err := os.Open("config/config-2.yaml")
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

		// switch res.Kind {
		// case "wireguard/v1":
		// 	var wgSpec WireguardSpec
		// 	err = yaml.Unmarshal(specData, &wgSpec)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	res.Spec = wgSpec
		// case "squid/v1":
		// 	var squidSpec SquidSpec
		// 	err = yaml.Unmarshal(specData, &squidSpec)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	res.Spec = squidSpec
		// default:
		// 	return nil, fmt.Errorf("unknown kind: %s", res.Kind)
		// }

		resources = append(resources, &res)
	}

	return resources, nil
}
