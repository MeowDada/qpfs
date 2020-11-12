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
		rmCmd,
	},
}

var dirFlag = &cli.StringFlag{
	Name:    "dir",
	Aliases: []string{"d"},
	Usage:   "`DIR` to the local datastore",
	Value:   defaultOrbitDBPath(),
}

type driveLocation struct {
	Name string
	Hash string
}

func (d *driveLocation) GetPath(root string) string {
	return filepath.Join(root, d.Hash, d.Name)
}

func defaultOrbitDBPath() string {
	dir, _ := homedir.Dir()
	return filepath.Join(dir, ".orbitdb")
}
