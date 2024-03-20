package wireguard2

import "github.com/tenminutevpn/tenminutevpn-manager/network"

type Wireguard struct {
	PresharedKey string `yaml:"presharedkey,omitempty"`
	PrivateKey   string `yaml:"privatekey"`
	PublicKey    string `yaml:"publickey,omitempty,omitempty"`

	Address *network.Address `yaml:"address"`
	Port    int              `yaml:"port"`

	Peers []*Peer `yaml:"peers,omitempty"`
}
