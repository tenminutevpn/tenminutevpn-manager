package squid

import (
	"strings"
	"text/template"

	"github.com/tenminutevpn/tenminutevpn-manager/pkg/utils"
)

type Squid struct {
	Port int `yaml:"port"`
}

var squidTemplate *template.Template

func init() {
	tpl, err := utils.NewTemplate(templateFS, "templates/squid.conf.tpl")
	if err != nil {
		panic(err)
	}
	squidTemplate = tpl
}

func (squid *Squid) Template() *template.Template {
	return squidTemplate
}

type squidTemplateData struct {
	Port int
}

func makeSquidTemplateData(squid *Squid) *squidTemplateData {
	return &squidTemplateData{
		Port: squid.Port,
	}
}

func (squid *Squid) Render() string {
	var output strings.Builder
	squid.Template().Execute(&output, makeSquidTemplateData(squid))
	return output.String()
}

func (squid *Squid) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{
		"port": squid.Port,
	}, nil
}

func (squid *Squid) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type S Squid
	s := (*S)(squid)
	if err := unmarshal(&s); err != nil {
		return err
	}
	*squid = Squid(*s)
	return nil
}
