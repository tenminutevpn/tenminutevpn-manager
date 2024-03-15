package wireguard

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

type Wireguard struct {
	NetworkInterface   string
	WireguardInterface string

	IP    net.IP
	IPNet *net.IPNet
	Port  int

	PrivateKey string
	PublicKey  string
}

func NewWireguard(wireguardInterface string, networkInterface string, ip net.IP, ipNet *net.IPNet, port int) *Wireguard {
	return &Wireguard{
		NetworkInterface:   networkInterface,
		WireguardInterface: wireguardInterface,
		IP:                 ip,
		IPNet:              ipNet,
		Port:               port,
	}
}

func (wg *Wireguard) SetPrivateKey(privateKey string) error {
	wg.PrivateKey = privateKey
	publickey, err := GenPublicKey(privateKey)
	if err != nil {
		return err
	}
	wg.PublicKey = publickey
	return nil
}

func (wg *Wireguard) ToServerConfig() *serverConfig {
	addressMask, _ := wg.IPNet.Mask.Size()
	address := fmt.Sprintf("%s/%d", wg.IP.String(), addressMask)
	return makeServerConfig(
		wg.WireguardInterface,
		address,
		wg.PrivateKey,
		fmt.Sprintf("%d", wg.Port),
		wg.NetworkInterface,
	)
}

func GenKey() (string, error) {
	cmd := exec.Command("wg", "genkey")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to generate private key: %w", err)
	}
	privkey := strings.TrimSpace(string(out))
	return privkey, nil
}

func GenPublicKey(privkey string) (string, error) {
	cmd := exec.Command("wg", "pubkey")
	cmd.Stdin = strings.NewReader(privkey)
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to generate public key: %w", err)
	}
	pubkey := strings.TrimSpace(string(out))
	return pubkey, nil
}
