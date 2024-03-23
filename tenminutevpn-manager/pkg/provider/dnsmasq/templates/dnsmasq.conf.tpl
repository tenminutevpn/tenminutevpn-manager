interface={{ .Interface }}
bind-interfaces

{{- range .Server }}
server={{ . }}
{{- end }}

domain-needed
bogus-priv

cache-size=1000
