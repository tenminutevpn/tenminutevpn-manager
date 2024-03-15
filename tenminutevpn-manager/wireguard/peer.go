package wireguard

import (
	"strings"
	"text/template"
)

const peerTemplate = `[Peer]
PublicKey = {{ .Wireguard.KeyPair.PublicKey }}
AllowedIPs = {{ range .AllowedIPs }}{{ .String }}, {{ end }}
{{ if ne .Wireguard.Port 0 }}Endpoint = {{ .Wireguard.GetPublicIPv4 }}:{{ .Wireguard.Port }}{{ end }}
{{ if ne .PersistentKeepalive 0 }}PersistentKeepalive = {{ .PersistentKeepalive }}{{ end }}
`

type WireguardPeer struct {
	Wireguard           *Wireguard
	AllowedIPs          []*Address
	PersistentKeepalive int
}

func NewWireguardPeer(wg *Wireguard, addrs []string, persistentKeepalive int) (*WireguardPeer, error) {
	allowedIPs := make([]*Address, 0)

	for _, addr := range addrs {
		a, err := NewAddressFromString(addr)
		if err != nil {
			return nil, err
		}
		allowedIPs = append(allowedIPs, a)
	}

	return &WireguardPeer{
		Wireguard:           wg,
		AllowedIPs:          allowedIPs,
		PersistentKeepalive: persistentKeepalive,
	}, nil
}

func (p *WireguardPeer) Render() string {
	tpl := template.Must(template.New("peerConfig").Parse(peerTemplate))
	var output strings.Builder
	tpl.Execute(&output, p)
	return output.String()
}
