package wireguard2

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/tenminutevpn/tenminutevpn-manager/network"
	"github.com/tenminutevpn/tenminutevpn-manager/utils"
)

type WireGuard struct {
	PresharedKey *Key `yaml:"presharedkey,omitempty"`
	PrivateKey   *Key `yaml:"privatekey"`
	PublicKey    *Key `yaml:"publickey,omitempty,omitempty"`

	Address *network.Address `yaml:"address"`
	Port    int              `yaml:"port"`

	Peers []*Peer `yaml:"peers,omitempty"`
}

var wireguardTemplate *template.Template

func init() {
	tpl, err := utils.NewTemplate(templateFS, "templates/wireguard.conf.tpl")
	if err != nil {
		panic(err)
	}
	wireguardTemplate = tpl
}

func (wireguard *WireGuard) Template() *template.Template {
	return wireguardTemplate
}

type wireguardTemplateData struct {
	PresharedKey string
	PrivateKey   string
	PublicKey    string

	DNS              string // TODO: This should be a list of DNS servers
	Name             string // TODO: This should be the name of the Wireguard interface
	NetworkInterface string // TODO: This should be the name of the network interface

	Address string
	Port    int

	Peers []*Peer
}

func makeWireguardTemplateData(wireguard *WireGuard) *wireguardTemplateData {
	presharedKey := ""
	if wireguard.PresharedKey != nil {
		presharedKey = wireguard.PresharedKey.String()
	}

	privateKey := ""
	if wireguard.PrivateKey != nil {
		privateKey = wireguard.PrivateKey.String()
	}

	publicKey := ""
	if wireguard.PublicKey != nil {
		publicKey = wireguard.PublicKey.String()
	}

	return &wireguardTemplateData{
		PresharedKey: presharedKey,
		PrivateKey:   privateKey,
		PublicKey:    publicKey,

		Address: wireguard.Address.String(),
		Port:    wireguard.Port,

		Peers: wireguard.Peers,
	}
}

func (wireguard *WireGuard) Render() string {
	var output strings.Builder
	wireguard.Template().Execute(&output, makeWireguardTemplateData(wireguard))
	return output.String()
}

func (wireguard *WireGuard) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type W WireGuard
	wg := (*W)(wireguard)
	if err := unmarshal(wg); err != nil {
		return err
	}

	if wg.PrivateKey == nil {
		return fmt.Errorf("private key is required")
	}

	if wg.PublicKey == nil {
		k := wg.PrivateKey.PublicKey()
		wg.PublicKey = &k
	}

	*wireguard = WireGuard(*wg)
	return nil
}
