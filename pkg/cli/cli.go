package cli

import (
	"strings"
)

var defaultHelpTemplate = `COMMANDS:
{{range .Commands}}	{{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}`

func indent(spaces int, v string) string {
	pad := strings.Repeat(" ", spaces)
	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}

func nindent(spaces int, v string) string {
	return "\n" + indent(spaces, v)
}
