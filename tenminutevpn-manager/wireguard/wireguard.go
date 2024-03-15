package wireguard

import (
	"fmt"
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
	serverConfig := wg.GetConfig()
	return serverConfig.Write(filename)
}

func (wg *Wireguard) ToPeer(allowedIPs *Address) *WireguardPeer {
	return &WireguardPeer{
		Wireguard:  wg,
		AllowedIPs: allowedIPs,
	}
}

func (wg *Wireguard) AddPeer(client *Wireguard) error {
	server := wg

	clientPeer, err := NewWireguardPeer(client, client.Address.String())
	if err != nil {
		return fmt.Errorf("failed to create peer (server -> client): %w", err)
	}
	server.Peers = append(server.Peers, clientPeer)

	serverPeer, err := NewWireguardPeer(server, server.Address.String())
	if err != nil {
		return fmt.Errorf("failed to create peer (client -> server): %w", err)
	}
	client.Peers = append(client.Peers, serverPeer)

	return nil
}
