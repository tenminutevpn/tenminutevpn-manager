package squid

import (
	"fmt"

	"github.com/tenminutevpn/tenminutevpn-manager/resource"
	"github.com/tenminutevpn/tenminutevpn-manager/systemd"
	"github.com/tenminutevpn/tenminutevpn-manager/utils"
)

type Resource struct {
	Kind     string            `yaml:"kind"`
	Metadata resource.Metadata `yaml:"metadata"`
	Spec     *Squid            `yaml:"spec"`
}

type Options struct {
	ConfigDir string
}

func (r *Resource) Options() *Options {
	configDir := "/etc/squid"
	if r.Metadata.Annotations["tenminutevpn.com/config-dir"] != "" {
		configDir = r.Metadata.Annotations["tenminutevpn.com/config-dir"]
	}

	if ok := utils.PathExists(configDir); !ok {
		if err := utils.CreateDir(configDir); err != nil {
			return nil
		}
	}

	return &Options{
		ConfigDir: configDir,
	}
}

func (r *Resource) Create() error {
	config := r.Spec.Render()
	configPath := fmt.Sprintf("%s/squid.conf", r.Options().ConfigDir)
	if err := utils.WriteToFile(configPath, 0600, config); err != nil {
		return fmt.Errorf("failed to write squid config to file: %w", err)
	}

	return nil
}

func (r *Resource) Service() *systemd.Service {
	return systemd.NewService("squid")
}
