package wireguard

import "github.com/tenminutevpn/tenminutevpn-manager/resource"

type WireguardResource struct {
	resource.Resource

	Spec *Wireguard `yaml:"spec"`
}
