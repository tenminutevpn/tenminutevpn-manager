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
	NetworkInterface  string
	WireguardIterface string

	Address net.IPMask
	Port    int

	PrivateKey string
	PublicKey  string
}

func NewWireguard(networkInterface string, wireguardInterface string, address net.IPMask, port int) *Wireguard {
	return &Wireguard{
		NetworkInterface:  networkInterface,
		WireguardIterface: wireguardInterface,
		Address:           address,
		Port:              port,
	}
}

func (w *Wireguard) SetPrivateKey(privateKey string) error {
	w.PrivateKey = privateKey
	publickey, err := GenPublicKey(privateKey)
	if err != nil {
		return err
	}
	w.PublicKey = publickey
	return nil
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

func GenKeypair() (string, string, error) {
	privkey, err := GenKey()
	if err != nil {
		return "", "", err
	}

	pubkey, err := GenPublicKey(privkey)
	if err != nil {
		return "", "", err
	}

	return privkey, pubkey, nil
}

func WriteKeypair(folder string, privateKey string, publicKey string) error {
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

const serverConfigTemplate = `[Interface]
Address = {{ .Address }}
PrivateKey = {{ .PrivateKey }}
ListenPort = {{ .ListenPort }}
PostUp = iptables -A FORWARD -i {{ .WireguardInterface }} -j ACCEPT; iptables -t nat -A POSTROUTING -o {{ .Interface }} -j MASQUERADE
PostDown = iptables -D FORWARD -i {{ .WireguardInterface }} -j ACCEPT; iptables -t nat -D POSTROUTING -o {{ .Interface }} -j MASQUERADE
`

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
