[Interface]
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
