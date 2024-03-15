package wireguard

import (
	"strings"
	"text/template"
)

const peerConfigTemplate = `[Peer]
PublicKey = {{ .PublicKey }}
AllowedIPs = {{ .AllowedIPs }}
`

type peerConfig struct {
	PublicKey  string
	AllowedIPs string
}

func makePeerConfig(publicKey, allowedIPs string) *peerConfig {
	return &peerConfig{
		PublicKey:  publicKey,
		AllowedIPs: allowedIPs,
	}
}

func (cfg *peerConfig) Render() string {
	tpl := template.Must(template.New("peerConfig").Parse(peerConfigTemplate))
	var output strings.Builder
	tpl.Execute(&output, cfg)
	return output.String()
}

const configTemplate = `[Interface]
# Name = {{ .Name }}
PrivateKey = {{ .PrivateKey }}
Address = {{ .Address }}

{{- if ne .ListenPort "0" }}
ListenPort = {{ .ListenPort }}
PostUp = iptables -A FORWARD -i {{ .Name }} -j ACCEPT; iptables -t nat -A POSTROUTING -o {{ .NetworkInterface }} -j MASQUERADE
PostDown = iptables -D FORWARD -i {{ .Name }} -j ACCEPT; iptables -t nat -D POSTROUTING -o {{ .NetworkInterface }} -j MASQUERADE
{{- end }}

{{- range .Peers }}
{{ .Render }}
{{- end }}
`

type wireguardConfig struct {
	Name             string
	Address          string
	PrivateKey       string
	ListenPort       string
	NetworkInterface string
	Peers            []peerConfig
}

func makeWireguardConfig(name, address, privateKey, listenPort, networkInterface string, peers []peerConfig) *wireguardConfig {
	return &wireguardConfig{
		Name:             name,
		Address:          address,
		PrivateKey:       privateKey,
		ListenPort:       listenPort,
		NetworkInterface: networkInterface,
		Peers:            peers,
	}
}

func (cfg *wireguardConfig) Render() string {
	tpl := template.Must(template.New("serverConfig").Parse(configTemplate))
	var output strings.Builder
	tpl.Execute(&output, cfg)
	return output.String()
}

func (cfg *wireguardConfig) Write(filename string) error {
	data := cfg.Render()
	return writeToFile(filename, 0600, data)
}
