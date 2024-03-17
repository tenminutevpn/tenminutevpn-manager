package wireguard

import (
	"embed"
	"fmt"
	"io/fs"
	"text/template"
)

//go:embed templates
var templates embed.FS

func getTemplate(name string) (*template.Template, error) {
	tpl, err := fs.ReadFile(templates, fmt.Sprintf("templates/%s.tpl", name))
	if err != nil {
		return nil, fmt.Errorf("failed to read template %s: %w", name, err)
	}
	return template.New(name).Parse(string(tpl))
}
