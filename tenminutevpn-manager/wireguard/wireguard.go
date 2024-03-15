package wireguard

import (
	"fmt"
)

type Wireguard struct {
	Name    string
	KeyPair *KeyPair

	NetworkInterface string

	Address *Address
	Port    int
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
	)
}

func (wg *Wireguard) WriteConfig(filename string) error {
	serverConfig := wg.GetConfig()
	return serverConfig.Write(filename)
}
