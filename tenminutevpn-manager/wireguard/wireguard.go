package wireguard

import (
	"fmt"
	"net"
)

type Wireguard struct {
	NetworkInterface   string
	WireguardInterface string

	IP    net.IP
	IPNet *net.IPNet
	Port  int

	KeyPair *KeyPair
}

func NewWireguard(wireguardInterface string, networkInterface string, ip net.IP, ipNet *net.IPNet, port int) (*Wireguard, error) {
	privkey, err := GeneratePrivateKey()
	if err != nil {
		return nil, err
	}

	keyPair, err := NewKeyPair(privkey)
	if err != nil {
		return nil, err
	}

	return &Wireguard{
		NetworkInterface:   networkInterface,
		WireguardInterface: wireguardInterface,

		IP:    ip,
		IPNet: ipNet,
		Port:  port,

		KeyPair: keyPair,
	}, nil
}

func (wg *Wireguard) GetServerConfig() *serverConfig {
	addressMask, _ := wg.IPNet.Mask.Size()
	address := fmt.Sprintf("%s/%d", wg.IP.String(), addressMask)
	return makeServerConfig(
		wg.WireguardInterface,
		address,
		wg.KeyPair.PrivateKey,
		fmt.Sprintf("%d", wg.Port),
		wg.NetworkInterface,
	)
}

func (wg *Wireguard) WriteServerConfig(filename string) error {
	serverConfig := wg.GetServerConfig()
	return serverConfig.Write(filename)
}
