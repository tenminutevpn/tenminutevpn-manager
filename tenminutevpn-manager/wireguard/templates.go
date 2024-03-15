package wireguard

const serverConfigTemplate = `[Interface]
Address = {{ .Address }}/24
PrivateKey = {{ .PrivateKey }}
ListenPort = {{ .Port }}
PostUp = iptables -A FORWARD -i {{ .WireguardInterface }} -j ACCEPT; iptables -t nat -A POSTROUTING -o {{ .NetworkInterface }} -j MASQUERADE
PostDown = iptables -D FORWARD -i {{ .WireguardInterface }} -j ACCEPT; iptables -t nat -D POSTROUTING -o {{ .NetworkInterface }} -j MASQUERADE
`
