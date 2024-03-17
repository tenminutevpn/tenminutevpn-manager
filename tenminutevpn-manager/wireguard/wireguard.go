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

	Peers []*Peer
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

func (wg *Wireguard) GetPublicIPv4() (string, error) {
	ip, err := network.GetPublicIPv4()
	if err != nil {
		return "", fmt.Errorf("failed to get public IPv4: %w", err)
	}
	return ip.String(), nil
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

func (server *Wireguard) AddPeer(client *Wireguard) error {
	peer := &Peer{
		PublicKey:  client.KeyPair.PublicKey,
		AllowedIPs: []*Address{client.Address},
	}
	server.Peers = append(server.Peers, peer)

	allowedIPv4, err := NewAddressFromString("0.0.0.0/0")
	if err != nil {
		return fmt.Errorf("failed to create allowed IPv4: %w", err)
	}

	allowedIPv6, err := NewAddressFromString("::/0")
	if err != nil {
		return fmt.Errorf("failed to create allowed IPv6: %w", err)
	}

	endpointIPv4, err := server.GetPublicIPv4()
	if err != nil {
		return fmt.Errorf("failed to get public IPv4: %w", err)
	}

	peer = &Peer{
		PublicKey:  server.KeyPair.PublicKey,
		AllowedIPs: []*Address{allowedIPv4, allowedIPv6},
		Endpoint:   fmt.Sprintf("%s:%d", endpointIPv4, server.Port),
	}
	client.Peers = append(client.Peers, peer)

	return nil

	// server := wg

	// clientAllowedIPs := []string{client.Address.String()}
	// clientPeer, err := NewWireguardPeer(client, clientAllowedIPs, 0)
	// if err != nil {
	// 	return fmt.Errorf("failed to create peer (server -> client): %w", err)
	// }
	// server.Peers = append(server.Peers, clientPeer)

	// serverAllowedIPs := []string{"::/0", "0.0.0.0/0"}
	// serverPeer, err := NewWireguardPeer(server, serverAllowedIPs, 25)
	// if err != nil {
	// 	return fmt.Errorf("failed to create peer (client -> server): %w", err)
	// }
	// client.Peers = append(client.Peers, serverPeer)

	return nil
}
