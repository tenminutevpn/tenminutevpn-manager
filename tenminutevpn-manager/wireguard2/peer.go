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

func makePeerTemplateData(p *Peer) *peerTemplateData {
	allowedIPs := make([]string, 0, len(p.AllowedIPs))
	for _, allowedIP := range p.AllowedIPs {
		allowedIPs = append(allowedIPs, allowedIP.String())
	}

	return &peerTemplateData{
		PublicKey:           p.PublicKey,
		AllowedIPs:          strings.Join(allowedIPs, ", "),
		Endpoint:            p.Endpoint.String(),
		PersistentKeepalive: p.PersistentKeepalive,
		PresharedKey:        p.PresharedKey,
	}
}

func (peer *Peer) Render() string {
	var output strings.Builder
	peer.Template().Execute(&output, makePeerTemplateData(peer))
	return output.String()

}
