package wireguard2

import (
	"fmt"
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

	AllowedIPs []network.Address `yaml:"allowedips"`
	Endpoint   *network.Endpoint `yaml:"endpoint,omitempty"`

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

	if len(peer.AllowedIPs) == 0 {
		return fmt.Errorf("allowedips is required")
	}

	if len(peer.AllowedIPs) > 1 {
		return fmt.Errorf("only one allowedip is allowed")
	}

	if peer.PublicKey == nil {
		if peer.PrivateKey == nil {
			privatekey, err := NewKey()
			if err != nil {
				return fmt.Errorf("failed to generate private key: %w", err)
			}
			peer.PrivateKey = &privatekey
		}
		publickey := peer.PrivateKey.PublicKey()
		peer.PublicKey = &publickey

	} else {
		if peer.PrivateKey != nil {
			if peer.PublicKey.String() != peer.PrivateKey.PublicKey().String() {
				return fmt.Errorf("public key does not match private key")
			}
		}
	}

	return nil
}
