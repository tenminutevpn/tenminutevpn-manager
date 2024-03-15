package wireguard

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"
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

func WriteKeypair(folder string, privateKey string, publicKey string) error {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err := os.Mkdir(folder, 0700)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	privkeyPath := path.Join(folder, "privkey")
	privkeyFile, err := os.Create(privkeyPath)
	if err != nil {
		return fmt.Errorf("failed to create private key file: %w", err)
	}
	defer privkeyFile.Close()

	_, err = privkeyFile.WriteString(privateKey)
	if err != nil {
		return fmt.Errorf("failed to write private key to file: %w", err)
	}

	pubkeyPath := path.Join(folder, "pubkey")
	pubkeyFile, err := os.Create(pubkeyPath)
	if err != nil {
		return fmt.Errorf("failed to create public key file: %w", err)
	}
	defer pubkeyFile.Close()

	_, err = pubkeyFile.WriteString(publicKey)
	if err != nil {
		return fmt.Errorf("failed to write public key to file: %w", err)
	}

	return nil
}

func GenServerConfig(iface string, privateKey string) string {
	tpl := template.Must(template.New("serverConfig").Parse(serverConfigTemplate))
	data := struct {
		Address            string
		PrivateKey         string
		ListenPort         string
		WireguardInterface string
		Interface          string
	}{
		Address:            "100.96.0.1/24",
		PrivateKey:         privateKey,
		ListenPort:         "51820",
		WireguardInterface: "wg0",
		Interface:          iface,
	}

	var config strings.Builder
	tpl.Execute(&config, data)
	return config.String()
}
