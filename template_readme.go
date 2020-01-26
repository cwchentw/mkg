package main

const templateREADME = `# {{.Prog}}

{{.Brief}}

## Copyright

Copyright (c) {{.Year}} {{.Author}}{{ if .License -}}. Licensed under {{ .License }}.{{- end}}`
