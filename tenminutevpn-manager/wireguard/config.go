package wireguard

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

const configTemplate = `[Interface]
# Name = {{ .Name }}
PrivateKey = {{ .PrivateKey }}
Address = {{ .Address }}

{{- if ne .ListenPort "0" }}
ListenPort = {{ .ListenPort }}
PostUp = iptables -A FORWARD -i {{ .Name }} -j ACCEPT; iptables -t nat -A POSTROUTING -o {{ .NetworkInterface }} -j MASQUERADE
PostDown = iptables -D FORWARD -i {{ .Name }} -j ACCEPT; iptables -t nat -D POSTROUTING -o {{ .NetworkInterface }} -j MASQUERADE
{{- end }}
`

type wireguardConfig struct {
	Name             string
	Address          string
	PrivateKey       string
	ListenPort       string
	NetworkInterface string
}

func makeWireguardConfig(name, address, privateKey, listenPort, networkInterface string) *wireguardConfig {
	return &wireguardConfig{
		Name:             name,
		Address:          address,
		PrivateKey:       privateKey,
		ListenPort:       listenPort,
		NetworkInterface: networkInterface,
	}
}

func (cfg *wireguardConfig) Render() string {
	tpl := template.Must(template.New("serverConfig").Parse(configTemplate))
	var output strings.Builder
	tpl.Execute(&output, cfg)
	return output.String()
}

func (cfg *wireguardConfig) Write(filename string) error {
	if filename == "" {
		return fmt.Errorf("filename is empty")
	}

	if _, err := os.Stat(filename); err == nil {
		return fmt.Errorf("file already exists: %s", filename)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	err = file.Chmod(0600)
	if err != nil {
		return fmt.Errorf("failed to set file permissions: %w", err)
	}

	_, err = file.WriteString(cfg.Render())
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
