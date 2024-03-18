[Peer]
PublicKey = {{ .PublicKey }}
AllowedIPs = {{ .AllowedIPs }}
{{ if ne .Endpoint "" }}Endpoint = {{ .Endpoint }}{{ end }}
{{ if ne .PersistentKeepalive 0 }}PersistentKeepalive = {{ .PersistentKeepalive }}{{ end }}
{{ if ne .PresharedKey "" }}PresharedKey = {{ .PresharedKey }}{{ end }}
