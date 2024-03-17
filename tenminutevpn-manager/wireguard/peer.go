package wireguard

import (
	"strings"
)

type Peer struct {
	PublicKey           string
	AllowedIPs          []*Address
	Endpoint            string
	PersistentKeepalive int
	PresharedKey        string
}

func (p *Peer) toTemplateData() *templatePeerData {
	allowedIPs := make([]string, 0, len(p.AllowedIPs))
	for _, allowedIP := range p.AllowedIPs {
		allowedIPs = append(allowedIPs, allowedIP.String())
	}

	return &templatePeerData{
		PublicKey:           p.PublicKey,
		AllowedIPs:          strings.Join(allowedIPs, ", "),
		Endpoint:            p.Endpoint,
		PersistentKeepalive: p.PersistentKeepalive,
		PresharedKey:        p.PresharedKey,
	}
}

func (p *Peer) Render() string {
	var output strings.Builder
	templatePeer.Execute(&output, p.toTemplateData())
	return output.String()
}
