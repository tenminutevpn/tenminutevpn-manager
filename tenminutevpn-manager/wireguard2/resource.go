package wireguard2

import "github.com/tenminutevpn/tenminutevpn-manager/resource"

type Resource struct {
	Kind     string            `yaml:"kind"`
	Metadata resource.Metadata `yaml:"metadata"`
	Spec     *Wireguard        `yaml:"spec"`
}
