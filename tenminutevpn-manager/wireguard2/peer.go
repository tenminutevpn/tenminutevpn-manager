package wireguard2

import "github.com/tenminutevpn/tenminutevpn-manager/network"

type Peer struct {
	PresharedKey string `yaml:"presharedKey,omitempty"`
	PrivateKey   string `yaml:"privateKey,omitempty"`
	PublicKey    string `yaml:"publicKey"`

	AllowedIPs []*network.Address `yaml:"allowedips"`
	Endpoint   *network.Address   `yaml:"endpoint,omitempty"`

	PersistentKeepalive int `yaml:"persistentkeepalive,omitempty"`
}
