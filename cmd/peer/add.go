package peer

import (
	"fmt"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
)

func defaultRepoPath() string {
	dir, _ := homedir.Dir()
	return filepath.Join(dir, ".ipfs")
}

var addCmd = &cli.Command{
	Name:  "add",
	Usage: "Add the peer address to the bootstrap persistent",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "dir",
			Aliases: []string{"d"},
			Usage:   "Directory to the IPFS repo",
			Value:   defaultRepoPath(),
		},
	},
	Action: func(c *cli.Context) error {
		return fmt.Errorf("seems unnecessary now")
	},
}
