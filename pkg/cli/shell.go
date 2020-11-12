package cli

import (
	"context"
	"io"
	"os"
	"strings"
	"text/template"
)

// Shell is the instance of the cli tool.
type Shell struct {
	Commands []*Command
	Usage    string
	Extra    map[string]interface{}
}

// Set sets the Extra field of the shell. It adopts last-one-wins fashion to
// handle duplicate names.
func (s *Shell) Set(k string, v interface{}) {
	if s.Extra == nil {
		s.Extra = make(map[string]interface{})
	}
	s.Extra[k] = v
}

// Get gets the value from Extra field of the shell.
func (s *Shell) Get(key string) (interface{}, bool) {
	v, ok := s.Extra[key]
	return v, ok
}

// Run runs the cli program.
func (s *Shell) Run(ctx context.Context, args []string) error {
	cmds := map[string]*Command{}
	for _, cmd := range s.Commands {
		cmdNames := append(cmd.Aliases, cmd.Name)
		for i := range cmdNames {
			cmds[cmdNames[i]] = cmd
		}
	}

	if len(args) == 1 && args[0] == "" {
		s.WriteUsage(os.Stdout)
		return nil
	}

	cmd := args[0]
	var cleanedArgs []string
	if len(args) > 1 {
		cleanedArgs = args[1:]
	}

	c, ok := cmds[cmd]
	if ok {
		return c.Do(ctx, s, cleanedArgs)
	}
	return s.hasNoSuchCommand(ctx, cmd)
}

// WriteUsage writes help usage text to the given writer.
func (s *Shell) WriteUsage(w io.Writer) error {
	funcMap := template.FuncMap{
		"join":    strings.Join,
		"indent":  indent,
		"nindent": nindent,
		"trim":    strings.TrimSpace,
	}
	t := template.Must(template.New("help").Funcs(funcMap).Parse(defaultHelpTemplate))
	return t.Execute(w, s)
}

func (s *Shell) hasNoSuchCommand(ctx context.Context, cmd string) error {
	return nil
}

func (s *Shell) printUsage() {

}
