package main

const templateREADME = `# {{.Prog}}

{{.Brief}}

## Author

{{.Year}}, {{.Author}}

{{ if .License -}}
## Copyright

{{ .License }}{{- end}}`
