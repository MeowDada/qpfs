package cli

import (
	"context"
	"sort"
)

// Command represents a cli command.
type Command struct {
	Name      string
	Aliases   []string
	Usage     string
	UsageText string
	Before    func(c *Context) error
	Action    func(c *Context) error
}

// Do process the input data.
func (cmd *Command) Do(ctx context.Context, shell *Shell, args []string) (err error) {
	c := &Context{ctx: ctx, shell: shell, args: args}
	if cmd.Before != nil {
		err = cmd.Before(c)
	}
	if err != nil {
		return err
	}
	if cmd.Action != nil {
		err = cmd.Action(c)
	}
	return err
}

// Names return the names including short names and aliases.
func (cmd *Command) Names() []string {
	return append([]string{cmd.Name}, cmd.Aliases...)
}

// RegisterCommand register commands to a shell.
func RegisterCommand(s *Shell, cmds ...*Command) {
	s.Commands = append(s.Commands, cmds...)
	sort.Slice(s.Commands, func(i, j int) bool {
		return s.Commands[i].Name < s.Commands[j].Name
	})
}
