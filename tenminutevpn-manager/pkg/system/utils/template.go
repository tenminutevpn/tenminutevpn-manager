package utils

import (
	"embed"
	"fmt"
	"text/template"
)

func NewTemplate(fs embed.FS, filename string) (*template.Template, error) {
	tpl, err := fs.ReadFile(fmt.Sprintf("%s", filename))
	if err != nil {
		return nil, fmt.Errorf("failed to read template %s: %w", filename, err)
	}
	return template.Must(template.New(filename).Parse(string(tpl))), nil
}
