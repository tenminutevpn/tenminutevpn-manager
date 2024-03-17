package wireguard

import (
	"fmt"
	"strings"

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

	endpointIPv4, err := network.GetPublicIPv4()
	if err != nil {
		return fmt.Errorf("failed to get public IPv4: %w", err)
	}

	peer = &Peer{
		PublicKey:  server.KeyPair.PublicKey,
		AllowedIPs: []*Address{allowedIPv4, allowedIPv6},
		Endpoint:   fmt.Sprintf("%s:%d", endpointIPv4.String(), server.Port),
	}
	client.Peers = append(client.Peers, peer)

	return nil
}

func (wg *Wireguard) toTemplateData() *templateWireguardData {
	return &templateWireguardData{
		Name:             wg.Name,
		PrivateKey:       wg.KeyPair.PrivateKey,
		Address:          wg.Address.String(),
		ListenPort:       wg.Port,
		NetworkInterface: wg.NetworkInterface,
		Peers:            wg.Peers,
	}
}

func (wg *Wireguard) Render() string {
	var output strings.Builder
	templateWireguard.Execute(&output, wg.toTemplateData())
	return output.String()
}

func (wg *Wireguard) Write(filename string) error {
	data := wg.Render()
	return writeToFile(filename, 0600, data)
}
