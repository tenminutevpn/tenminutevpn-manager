package wireguard

import (
	"embed"
	"fmt"
	"io/fs"
	"strings"
	"text/template"
)

//go:embed templates
var templates embed.FS

func getTemplate(name string) *template.Template {
	tpl, err := fs.ReadFile(templates, fmt.Sprintf("templates/%s.tpl", name))
	if err != nil {
		panic(fmt.Errorf("failed to read template %s: %w", name, err))
	}
	return template.Must(template.New(name).Parse(string(tpl)))
}

var (
	templatePeer      *template.Template
	templateWireguard *template.Template
)

func init() {
	templatePeer = getTemplate("peer.conf")
	templateWireguard = getTemplate("wireguard.conf")
}

type templatePeerData struct {
	PublicKey           string
	AllowedIPs          string
	Endpoint            string
	PersistentKeepalive int
	PresharedKey        string
}

func makeTemplatePeerData(p *Peer) *templatePeerData {
	allowedIPs := make([]string, 0, len(p.AllowedIPs))
	for _, allowedIP := range p.AllowedIPs {
		allowedIPs = append(allowedIPs, allowedIP.String())
	}

	return &templatePeerData{
		PublicKey:           p.PublicKey,
		AllowedIPs:          strings.Join(allowedIPs, ", "),
		Endpoint:            p.Endpoint,
		PersistentKeepalive: p.PersistentKeepalive,
		PresharedKey:        p.PresharedKey,
	}
}

type templateWireguardData struct {
	Name             string
	PrivateKey       string
	Address          string
	ListenPort       int
	NetworkInterface string
	DNS              string
	Peers            []*Peer
}
