package drive

import (
	"path/filepath"

	"github.com/meowdada/qpfs/cmd"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
)

func init() {
	cmd.Register(driveCmd)
}

var driveCmd = &cli.Command{
	Name:  "drive",
	Usage: "Operations to interact with local drive",
	Subcommands: []*cli.Command{
		lsCmd,
	},
}

func defaultOrbitDBPath() string {
	dir, _ := homedir.Dir()
	return filepath.Join(dir, ".orbitdb")
}
