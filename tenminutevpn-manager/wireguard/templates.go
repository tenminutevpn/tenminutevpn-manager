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
	templateServer *template.Template
	templateClient *template.Template
)

func init() {
	templateServer = getTemplate("server.conf")
	templateClient = getTemplate("client.conf")
}
