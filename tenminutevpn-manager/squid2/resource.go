package squid2

import (
	"github.com/tenminutevpn/tenminutevpn-manager/resource"
	"github.com/tenminutevpn/tenminutevpn-manager/systemd"
)

type Resource struct {
	Kind     string            `yaml:"kind"`
	Metadata resource.Metadata `yaml:"metadata"`
	Spec     *Squid            `yaml:"spec"`
}

func (r *Resource) Create() error {
	return nil
}

func (r *Resource) Service() *systemd.Service {
	return systemd.NewService("squid")
}
