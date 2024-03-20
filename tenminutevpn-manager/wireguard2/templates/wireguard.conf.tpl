[Interface]
PrivateKey = {{ .PrivateKey }}
Address = {{ .Address }}
{{- if .DNS }}
DNS = {{ .DNS }}
{{- end }}
{{- if ne .Port 0 }}
ListenPort = {{ .Port }}
PostUp = iptables -A FORWARD -i {{ .Name }} -j ACCEPT; iptables -t nat -A POSTROUTING -o {{ .NetworkInterface }} -j MASQUERADE
PostDown = iptables -D FORWARD -i {{ .Name }} -j ACCEPT; iptables -t nat -D POSTROUTING -o {{ .NetworkInterface }} -j MASQUERADE
{{- end }}

{{- range .Peers }}
{{ .Render }}
{{- end }}
