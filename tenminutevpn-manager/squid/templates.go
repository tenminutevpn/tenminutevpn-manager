package squid

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"strings"
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
	templateSquid *template.Template
)

func init() {
	templateSquid = getTemplate("squid.conf")
}

type templateSquidData struct {
	Port int
}

func makeTemplateSquidData(s *Squid) *templateSquidData {
	return &templateSquidData{
		Port: s.Port,
	}
}

func (t *templateSquidData) Render() string {
	var output strings.Builder
	templateSquid.Execute(&output, t)
	return output.String()
}
