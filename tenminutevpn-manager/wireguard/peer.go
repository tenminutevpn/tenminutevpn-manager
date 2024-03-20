package wireguard

import "github.com/tenminutevpn/tenminutevpn-manager/network"

type Peer struct {
	PublicKey           string
	AllowedIPs          []*network.Address
	Endpoint            string
	PersistentKeepalive int
	PresharedKey        string
}

func (p *Peer) Render() string {
	return makeTemplatePeerData(p).Render()
}
