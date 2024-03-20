package wireguard

import "github.com/tenminutevpn/tenminutevpn-manager/network"

type Peer struct {
	PresharedKey string
	PrivateKey   string
	PublicKey    string

	AllowedIPs []*network.Address
	Endpoint   string

	PersistentKeepalive int
}

func (p *Peer) Render() string {
	return makeTemplatePeerData(p).Render()
}
