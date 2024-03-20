package wireguard2

import "github.com/tenminutevpn/tenminutevpn-manager/network"

type Wireguard struct {
	PrivateKey string `yaml:"privateKey"`
	PublicKey  string `yaml:"publicKey,omitempty"`

	Address *network.Address `yaml:"address"`
	Port    int              `yaml:"port"`

	Peers []*Peer `yaml:"peers"`
}
