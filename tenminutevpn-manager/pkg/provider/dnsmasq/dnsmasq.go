package dnsmasq

import (
	"strings"
	"text/template"

	"github.com/tenminutevpn/tenminutevpn-manager/pkg/system/utils"
)

type Dnsmasq struct {
	Interface string   `yaml:"interface"`
	Server    []string `yaml:"server"`
}

var tpl *template.Template

func init() {
	t, err := utils.NewTemplate(templateFS, "templates/dnsmasq.conf.tpl")
	if err != nil {
		panic(err)
	}
	tpl = t
}

func (d *Dnsmasq) Template() *template.Template {
	return tpl
}

type templateData struct {
	Interface string
	Server    []string
}

func makeTemplateData(d *Dnsmasq) *templateData {
	return &templateData{
		Interface: d.Interface,
		Server:    d.Server,
	}
}

func (dnsmasq *Dnsmasq) Render() string {
	var output strings.Builder
	dnsmasq.Template().Execute(&output, makeTemplateData(dnsmasq))
	return output.String()
}

func (dnsmasq *Dnsmasq) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"interface": dnsmasq.Interface,
		"server":    dnsmasq.Server,
	}, nil
}

func (dnsmasq *Dnsmasq) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type D Dnsmasq
	d := (*D)(dnsmasq)
	if err := unmarshal(&d); err != nil {
		return err
	}
	*dnsmasq = Dnsmasq(*d)
	return nil
}
