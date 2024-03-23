package dnsmasq

import (
	"fmt"

	"github.com/tenminutevpn/tenminutevpn-manager/pkg/resource"
	"github.com/tenminutevpn/tenminutevpn-manager/pkg/system/systemd"
	"github.com/tenminutevpn/tenminutevpn-manager/pkg/system/utils"
)

type Resource struct {
	Kind     string            `yaml:"kind"`
	Metadata resource.Metadata `yaml:"metadata"`
	Spec     *Dnsmasq          `yaml:"spec"`
}

type Options struct {
}

func (r *Resource) Options() *Options {
	return &Options{}
}

func (r *Resource) Create() error {
	config := r.Spec.Render()
	configPath := "/etc/dnsmasq.conf"
	if err := utils.WriteToFile(configPath, 0600, config); err != nil {
		return fmt.Errorf("failed to write dnsmasq config to file: %w", err)
	}
	return nil
}

func (r *Resource) Service() *systemd.Service {
	return systemd.NewService("dnsmasq")
}
