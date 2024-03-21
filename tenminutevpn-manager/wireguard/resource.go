package wireguard

import (
	"fmt"

	"github.com/tenminutevpn/tenminutevpn-manager/resource"
	"github.com/tenminutevpn/tenminutevpn-manager/systemd"
	"github.com/tenminutevpn/tenminutevpn-manager/utils"
)

type Resource struct {
	Kind     string            `yaml:"kind"`
	Metadata resource.Metadata `yaml:"metadata"`
	Spec     *WireGuard        `yaml:"spec"`
}

type Options struct {
	ConfigDir     string
	PeerConfigDir string
}

func (r *Resource) Options() *Options {
	configDir := fmt.Sprintf("/etc/wireguard/")
	if r.Metadata.Annotations["tenminutevpn.com/config-dir"] != "" {
		configDir = r.Metadata.Annotations["tenminutevpn.com/config-dir"]
	}

	if ok := utils.PathExists(configDir); !ok {
		utils.CreateDir(configDir)
	}

	peerConfigDir := fmt.Sprintf("/etc/wireguard/peers")
	if r.Metadata.Annotations["tenminutevpn.com/peer-config-dir"] != "" {
		peerConfigDir = r.Metadata.Annotations["tenminutevpn.com/peer-config-dir"]
	}

	if ok := utils.PathExists(peerConfigDir); !ok {
		utils.CreateDir(peerConfigDir)
	}

	return &Options{
		ConfigDir:     configDir,
		PeerConfigDir: peerConfigDir,
	}
}

func (r *Resource) Create() error {
	deviceConfig := r.Spec.Render()
	deviceConfigPath := fmt.Sprintf("%s/%s.conf", r.Options().ConfigDir, r.Metadata.Name)
	if err := utils.WriteToFile(deviceConfigPath, 0600, deviceConfig); err != nil {
		return fmt.Errorf("failed to write wireguard config to file: %w", err)
	}

	for peerID, peer := range r.Spec.Peers {
		peerConfig := r.Spec.PeerWireguard(peer).Render()
		peerConfigPath := fmt.Sprintf("%s/peer-%d.conf", r.Options().PeerConfigDir, peerID)
		if err := utils.WriteToFile(peerConfigPath, 0600, peerConfig); err != nil {
			return fmt.Errorf("failed to write wireguard peer config to file: %w", err)
		}
	}

	return nil
}

func (r *Resource) Service() *systemd.Service {
	return systemd.NewService(fmt.Sprintf("wg-quick@%s", r.Spec.Device))
}
