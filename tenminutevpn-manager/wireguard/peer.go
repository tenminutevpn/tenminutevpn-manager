package wireguard

import (
	"strings"
	"text/template"
)

const peerTemplate = `[Peer]
PublicKey = {{ .Wireguard.KeyPair.PublicKey }}
AllowedIPs = {{ .AllowedIPs }}
{{ if ne .Wireguard.Port 0 }}Endpoint = {{ .Wireguard.GetPublicIPv4 }}:{{ .Wireguard.Port }}{{ end }}
`

type WireguardPeer struct {
	Wireguard  *Wireguard
	AllowedIPs *Address
}

func NewWireguardPeer(wg *Wireguard, addr string) (*WireguardPeer, error) {
	allowedIPs, err := NewAddressFromString(addr)
	if err != nil {
		return nil, err
	}
	return &WireguardPeer{
		Wireguard:  wg,
		AllowedIPs: allowedIPs,
	}, nil
}

func (p *WireguardPeer) Render() string {
	tpl := template.Must(template.New("peerConfig").Parse(peerTemplate))
	var output strings.Builder
	tpl.Execute(&output, p)
	return output.String()
}
