[Interface]
PrivateKey = {{ .PrivateKey }}
Address = {{ .Address }}

{{- if .DNS }}
DNS = {{ .DNS }}
{{- end }}

{{- if ne .Port 0 }}
ListenPort = {{ .Port }}
{{- end }}

{{- if .Device }}
PostUp = iptables -A FORWARD -i {{ .Device }} -j ACCEPT; iptables -t nat -A POSTROUTING -o {{ .DeviceRoute }} -j MASQUERADE
PostDown = iptables -D FORWARD -i {{ .Device }} -j ACCEPT; iptables -t nat -D POSTROUTING -o {{ .DeviceRoute }} -j MASQUERADE
{{- end }}

{{- range .Peers }}
{{ .Render }}
{{- end }}
