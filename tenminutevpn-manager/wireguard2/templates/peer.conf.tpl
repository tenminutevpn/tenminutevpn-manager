
[Peer]
{{- if ne .PresharedKey "" }}
PresharedKey = {{ .PresharedKey }}
{{ end }}

PublicKey = {{ .PublicKey }}
AllowedIPs = {{ .AllowedIPs }}

{{- if ne .Endpoint "" }}
Endpoint = {{ .Endpoint }}
{{ end }}

{{- if ne .PersistentKeepalive 0 }}
PersistentKeepalive = {{ .PersistentKeepalive }}
{{ end }}
