package drive

import (
	"github.com/meowdada/qpfs/cmd"
	"github.com/urfave/cli/v2"
)

func init() {
	cmd.Register(driveCmd)
}

var driveCmd = &cli.Command{
	Name:        "drive",
	Usage:       "Perform drive operations",
	Description: "",
	Subcommands: []*cli.Command{
		lsCmd,
		newCmd,
		statCmd,
		rmCmd,
		openCmd,
	},
}
