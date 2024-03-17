package wireguard

import (
	"strings"
)

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
	var output strings.Builder
	templateClient.Execute(&output, p)
	return output.String()
}
