package wireguard

import (
	"fmt"

	"github.com/tenminutevpn/tenminutevpn-manager/network"
)

type Wireguard struct {
	Name    string
	KeyPair *KeyPair

	NetworkInterface string
	Address          *Address
	Port             int

	Peers []*WireguardPeer
}

func NewWireguard(name string, networkInterface string, addr string, port int) (*Wireguard, error) {
	address, err := NewAddressFromString(addr)
	if err != nil {
		return nil, err
	}

	privkey, err := GeneratePrivateKey()
	if err != nil {
		return nil, err
	}

	keyPair, err := NewKeyPair(privkey)
	if err != nil {
		return nil, err
	}

	return &Wireguard{
		Name:    name,
		KeyPair: keyPair,

		NetworkInterface: networkInterface,

		Address: address,
		Port:    port,
	}, nil
}

func (wg *Wireguard) GetPublicIPv4() string {
	ip, err := network.GetPublicIPv4()
	if err != nil {
		return ""
	}
	return ip.String()
}

func (wg *Wireguard) GetConfig() *wireguardConfig {
	return makeWireguardConfig(
		wg.Name,
		wg.Address.String(),
		wg.KeyPair.PrivateKey,
		fmt.Sprintf("%d", wg.Port),
		wg.NetworkInterface,
		wg.Peers,
	)
}

func (wg *Wireguard) WriteConfig(filename string) error {
	return wg.GetConfig().Write(filename)
}

func (wg *Wireguard) ToPeer(allowedIPs []*Address) *WireguardPeer {
	return &WireguardPeer{
		Wireguard:  wg,
		AllowedIPs: allowedIPs,
	}
}

func (wg *Wireguard) AddPeer(client *Wireguard) error {
	server := wg

	clientAllowedIPs := []string{client.Address.String()}
	clientPeer, err := NewWireguardPeer(client, clientAllowedIPs, 0)
	if err != nil {
		return fmt.Errorf("failed to create peer (server -> client): %w", err)
	}
	server.Peers = append(server.Peers, clientPeer)

	serverAllowedIPs := []string{"::/0", "0.0.0.0/0"}
	serverPeer, err := NewWireguardPeer(server, serverAllowedIPs, 25)
	if err != nil {
		return fmt.Errorf("failed to create peer (client -> server): %w", err)
	}
	client.Peers = append(client.Peers, serverPeer)

	return nil
}
