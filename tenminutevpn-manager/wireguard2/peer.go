package wireguard2

import (
	"strings"
	"text/template"

	"github.com/tenminutevpn/tenminutevpn-manager/network"
	"github.com/tenminutevpn/tenminutevpn-manager/utils"
)

type Peer struct {
	PresharedKey string `yaml:"presharedkey,omitempty"`
	PrivateKey   string `yaml:"privateey,omitempty"`
	PublicKey    string `yaml:"publickey"`

	AllowedIPs []*network.Address `yaml:"allowedips"`
	Endpoint   *network.Endpoint  `yaml:"endpoint,omitempty"`

	PersistentKeepalive int `yaml:"persistentkeepalive,omitempty"`
}

var peerTemplate *template.Template

func init() {
	tpl, err := utils.NewTemplate(templateFS, "templates/peer.conf.tpl")
	if err != nil {
		panic(err)
	}
	peerTemplate = tpl
}

func (peer *Peer) Template() *template.Template {
	return peerTemplate
}

type peerTemplateData struct {
	PresharedKey string
	PublicKey    string

	AllowedIPs string
	Endpoint   string

	PersistentKeepalive int
}

func makePeerTemplateData(peer *Peer) *peerTemplateData {
	allowedIPs := make([]string, 0, len(peer.AllowedIPs))
	for _, allowedIP := range peer.AllowedIPs {
		allowedIPs = append(allowedIPs, allowedIP.String())
	}

	endpoint := ""
	if peer.Endpoint != nil {
		endpoint = peer.Endpoint.String()
	}

	return &peerTemplateData{
		PublicKey:           peer.PublicKey,
		AllowedIPs:          strings.Join(allowedIPs, ", "),
		Endpoint:            endpoint,
		PersistentKeepalive: peer.PersistentKeepalive,
		PresharedKey:        peer.PresharedKey,
	}
}

func (peer *Peer) Render() string {
	var output strings.Builder
	peer.Template().Execute(&output, makePeerTemplateData(peer))
	return output.String()

}

func (peer *Peer) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type P Peer
	if err := unmarshal((*P)(peer)); err != nil {
		return err
	}

	return nil
}
