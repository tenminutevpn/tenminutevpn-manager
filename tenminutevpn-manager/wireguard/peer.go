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

func (p *Peer) Render() string {
	var output strings.Builder
	templatePeer.Execute(&output, makeTemplatePeerData(p))
	return output.String()
}
