package wireguard2

import "github.com/tenminutevpn/tenminutevpn-manager/network"

type Peer struct {
	PresharedKey string `yaml:"presharedkey,omitempty"`
	PrivateKey   string `yaml:"privateey,omitempty"`
	PublicKey    string `yaml:"publickey"`

	AllowedIPs []*network.Address `yaml:"allowedips"`
	Endpoint   *network.Address   `yaml:"endpoint,omitempty"`

	PersistentKeepalive int `yaml:"persistentkeepalive,omitempty"`
}
