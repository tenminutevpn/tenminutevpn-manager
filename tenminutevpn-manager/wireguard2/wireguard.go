package wireguard2

import (
	"strings"
	"text/template"

	"github.com/tenminutevpn/tenminutevpn-manager/network"
	"github.com/tenminutevpn/tenminutevpn-manager/utils"
)

type Wireguard struct {
	PresharedKey string `yaml:"presharedkey,omitempty"`
	PrivateKey   string `yaml:"privatekey"`
	PublicKey    string `yaml:"publickey,omitempty,omitempty"`

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

func (wireguard *Wireguard) Template() *template.Template {
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

func makeWireguardTemplateData(wireguard *Wireguard) *wireguardTemplateData {
	return &wireguardTemplateData{
		PresharedKey: wireguard.PresharedKey,
		PrivateKey:   wireguard.PrivateKey,
		PublicKey:    wireguard.PublicKey,

		Address: wireguard.Address.String(),
		Port:    wireguard.Port,

		Peers: wireguard.Peers,
	}
}

func (wireguard *Wireguard) Render() string {
	var output strings.Builder
	wireguard.Template().Execute(&output, makeWireguardTemplateData(wireguard))
	return output.String()
}
