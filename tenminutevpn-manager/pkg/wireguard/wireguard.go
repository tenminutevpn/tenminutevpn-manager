package wireguard

import (
	"fmt"
	"net"
	"strings"
	"text/template"

	"github.com/tenminutevpn/tenminutevpn-manager/pkg/network"
	"github.com/tenminutevpn/tenminutevpn-manager/pkg/utils"
)

type WireGuard struct {
	Device string `yaml:"device"`

	PresharedKey *Key `yaml:"presharedkey,omitempty"`
	PrivateKey   *Key `yaml:"privatekey"`
	PublicKey    *Key `yaml:"publickey,omitempty,omitempty"`

	Address *network.Address `yaml:"address"`
	Port    int              `yaml:"port"`

	Peers []*Peer   `yaml:"peers,omitempty"`
	DNS   []*net.IP `yaml:"dns,omitempty"`
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
	Device      string
	DeviceRoute string

	PresharedKey string
	PrivateKey   string
	PublicKey    string

	Address string
	Port    int

	Peers []*Peer
	DNS   string
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

	dns := make([]string, 0, len(wireguard.DNS))
	for _, ip := range wireguard.DNS {
		dns = append(dns, ip.String())
	}

	route, err := network.GetDefaultInterface()
	if err != nil {
		panic(err)
	}

	return &wireguardTemplateData{
		Device:      wireguard.Device,
		DeviceRoute: route,

		PresharedKey: presharedKey,
		PrivateKey:   privateKey,
		PublicKey:    publicKey,

		Address: wireguard.Address.String(),
		Port:    wireguard.Port,

		Peers: wireguard.Peers,
		DNS:   strings.Join(dns, ", "),
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
		privatekey, err := NewKey()
		if err != nil {
			return fmt.Errorf("failed to generate private key: %w", err)
		}
		wg.PrivateKey = &privatekey
	}

	if wg.PublicKey == nil {
		k := wg.PrivateKey.PublicKey()
		wg.PublicKey = &k
	} else {
		if wg.PublicKey.String() != wg.PrivateKey.PublicKey().String() {
			return fmt.Errorf("public key does not match private key")
		}
	}

	if wg.Port == 0 {
		wg.Port = 51820
	} else if wg.Port < 1 || wg.Port > 65535 {
		return fmt.Errorf("port must be between 1 and 65535")
	}

	*wireguard = WireGuard(*wg)
	return nil
}

func (wireguard *WireGuard) PeerWireguard(client *Peer) *WireGuard {
	ip, err := network.GetPublicIPv4()
	if err != nil {
		panic(err)
	}

	endpoint := network.NewEndpoint(ip, wireguard.Port)

	allowedIPv4, _ := network.NewAddressFromString("0.0.0.0/0")
	allowedIPv6, _ := network.NewAddressFromString("::/0")

	peer := &Peer{
		PresharedKey: wireguard.PresharedKey,
		PublicKey:    wireguard.PublicKey,
		Endpoint:     endpoint,
		AllowedIPs:   []network.Address{*allowedIPv4, *allowedIPv6},
	}

	return &WireGuard{
		PresharedKey: wireguard.PresharedKey,
		PrivateKey:   client.PrivateKey,
		PublicKey:    client.PublicKey,

		Address: &client.AllowedIPs[0],

		DNS:   wireguard.DNS,
		Peers: []*Peer{peer},
	}
}
