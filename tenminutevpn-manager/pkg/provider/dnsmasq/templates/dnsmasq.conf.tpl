interface={{ .Interface }}
bind-interfaces

{{- range .Server }}
server={{ . }}
{{- end }}

domain-needed
bogus-priv
dnssec

cache-size=1000
