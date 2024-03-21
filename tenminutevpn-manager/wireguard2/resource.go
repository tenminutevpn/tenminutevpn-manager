package wireguard2

import (
	"fmt"

	"github.com/tenminutevpn/tenminutevpn-manager/resource"
	"github.com/tenminutevpn/tenminutevpn-manager/utils"
)

type Resource struct {
	Kind     string            `yaml:"kind"`
	Metadata resource.Metadata `yaml:"metadata"`
	Spec     *WireGuard        `yaml:"spec"`
}

func (r *Resource) deviceConfigPath() string {
	path := fmt.Sprintf("/etc/wireguard/%s.conf", r.Metadata.Name)
	if r.Metadata.Annotations["tenminutevpn.com/wireguard"] != "" {
		path = r.Metadata.Annotations["tenminutevpn.com/wireguard"]
	}
	return path
}

func (r *Resource) peerConfigPath() string {
	path := fmt.Sprintf("/etc/wireguard/peers")
	if r.Metadata.Annotations["tenminutevpn.com/wireguard-peers"] != "" {
		path = r.Metadata.Annotations["tenminutevpn.com/wireguard-peers"]
	}

	if ok := utils.PathExists(path); !ok {
		utils.CreateDir(path)
	}

	return path
}

func (r *Resource) Process() error {
	err := utils.WriteToFile(r.deviceConfigPath(), 0600, r.Spec.Render())
	if err != nil {
		return fmt.Errorf("failed to write wireguard config to file: %w", err)
	}

	for peerID, peer := range r.Spec.Peers {
		peerConfigPath := fmt.Sprintf("%s/peer-%d.conf", r.peerConfigPath(), peerID)

		peerWireguard := r.Spec.PeerWireguard(peer)
		err := utils.WriteToFile(peerConfigPath, 0600, peerWireguard.Render())
		if err != nil {
			return fmt.Errorf("failed to write wireguard peer config to file: %w", err)
		}
	}

	return nil
}
