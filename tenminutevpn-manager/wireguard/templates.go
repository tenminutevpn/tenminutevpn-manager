package wireguard

import (
	"embed"
	"fmt"
	"io/fs"
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
	templatePeer   *template.Template
	templateServer *template.Template
)

func init() {
	templatePeer = getTemplate("peer.conf")
	templateServer = getTemplate("server.conf")
}

type templatePeerData struct {
	PublicKey           string
	AllowedIPs          string
	Endpoint            string
	PersistentKeepalive int
	PresharedKey        string
}

type templateServerData struct{}
