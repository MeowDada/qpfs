package util

import (
	"github.com/meowdada/qpfs/cmd"
	"github.com/urfave/cli/v2"
)

func init() {
	cmd.Register(utilCmd)
}

var utilCmd = &cli.Command{
	Name:    "util",
	Aliases: []string{"u"},
	Usage:   "Providing some utilities helper functions",
	Subcommands: []*cli.Command{
		genBlockCmd,
		genFileCmd,
	},
}
