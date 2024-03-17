package wireguard

import (
	"strings"
)

type wireguardConfig struct {
	Name             string
	Address          string
	PrivateKey       string
	ListenPort       string
	NetworkInterface string
	Peers            []*WireguardPeer
}

func makeWireguardConfig(name, address, privateKey, listenPort, networkInterface string, peers []*WireguardPeer) *wireguardConfig {
	return &wireguardConfig{
		Name:             name,
		Address:          address,
		PrivateKey:       privateKey,
		ListenPort:       listenPort,
		NetworkInterface: networkInterface,
		Peers:            peers,
	}
}

func (cfg *wireguardConfig) Render() string {
	tpl, err := getTemplate("server.conf")
	if err != nil {
		panic(err)
	}

	var output strings.Builder
	tpl.Execute(&output, cfg)
	return output.String()
}

func (cfg *wireguardConfig) Write(filename string) error {
	data := cfg.Render()
	return writeToFile(filename, 0600, data)
}