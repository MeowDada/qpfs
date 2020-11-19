package key

import (
	"github.com/meowdada/qpfs/cmd"
	"github.com/urfave/cli/v2"
)

func init() {
	cmd.Register(keyCmd)
}

var keyCmd = &cli.Command{
	Name:  "key",
	Usage: "Swarm key management",
	Subcommands: []*cli.Command{
		lsCmd,
		newCmd,
	},
}
