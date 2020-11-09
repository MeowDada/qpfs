package orbitdb

import (
	"github.com/meowdada/qpfs/cmd"
	"github.com/urfave/cli/v2"
)

func init() {
	cmd.Register(orbitdbCmd)
}

var orbitdbCmd = &cli.Command{
	Name:        "orbitdb",
	Usage:       "Perform orbitdb raw operations",
	Description: "Directly perform orbitdb raw operations.",
	Subcommands: []*cli.Command{
		newCmd,
		addCmd,
		stopCmd,
		listCmd,
	},
}
