package wireguard

type Peer struct {
	PublicKey           string
	AllowedIPs          []*Address
	Endpoint            string
	PersistentKeepalive int
	PresharedKey        string
}

func (p *Peer) Render() string {
	return makeTemplatePeerData(p).Render()
}
