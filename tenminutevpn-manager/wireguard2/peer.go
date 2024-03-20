package wireguard2

import (
	"fmt"
	"net"
	"strings"
	"text/template"
	"time"

	"github.com/tenminutevpn/tenminutevpn-manager/network"
	"github.com/tenminutevpn/tenminutevpn-manager/utils"
)

type Peer struct {
	PresharedKey *Key `yaml:"presharedkey,omitempty"`
	PrivateKey   *Key `yaml:"privateey,omitempty"`
	PublicKey    *Key `yaml:"publickey"`

	AllowedIPs []network.IPNet `yaml:"allowedips"`
	Endpoint   *net.UDPAddr    `yaml:"endpoint,omitempty"`

	PersistentKeepalive time.Duration `yaml:"persistentkeepalive,omitempty"`
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
	var presharedKey string
	if peer.PresharedKey == nil {
		presharedKey = ""
	} else {
		presharedKey = peer.PresharedKey.String()
	}

	allowedIPs := make([]string, 0, len(peer.AllowedIPs))
	for _, allowedIP := range peer.AllowedIPs {
		allowedIPs = append(allowedIPs, allowedIP.String())
	}

	endpoint := ""
	if peer.Endpoint != nil {
		endpoint = peer.Endpoint.String()
	}

	persistentKeepalive := int(peer.PersistentKeepalive / time.Second)

	return &peerTemplateData{
		PresharedKey:        presharedKey,
		PublicKey:           peer.PublicKey.String(),
		AllowedIPs:          strings.Join(allowedIPs, ", "),
		Endpoint:            endpoint,
		PersistentKeepalive: persistentKeepalive,
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

	if peer.PrivateKey != nil {
		publickey := peer.PrivateKey.PublicKey()
		if peer.PublicKey == nil {
			peer.PublicKey = &publickey
		} else if *peer.PublicKey != publickey {
			return fmt.Errorf("public key does not match private key")
		}
	}

	if peer.PublicKey == nil {
		return fmt.Errorf("public key is required")
	}

	return nil
}
