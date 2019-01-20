package main

const templateREADME = `# {{.Prog}}

{{.Brief}}

## Copyright

{{.Year}}, {{.Author}}{{ if .License -}}; {{ .License }}{{- end}}`
