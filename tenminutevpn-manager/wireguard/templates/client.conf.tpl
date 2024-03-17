[Peer]
PublicKey = {{ .Wireguard.KeyPair.PublicKey }}
AllowedIPs = {{ range .AllowedIPs }}{{ .String }}, {{ end }}
{{ if ne .Wireguard.Port 0 }}Endpoint = {{ .Wireguard.GetPublicIPv4 }}:{{ .Wireguard.Port }}{{ end }}
{{ if ne .PersistentKeepalive 0 }}PersistentKeepalive = {{ .PersistentKeepalive }}{{ end }}
