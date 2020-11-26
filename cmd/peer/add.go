package peer

import (
	"errors"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
)

func defaultRepoPath() string {
	dir, _ := homedir.Dir()
	return filepath.Join(dir, ".ipfs")
}

func defaultClusterAPIEndpoint() string {
	return "/ip4/127.0.0.1/tcp/9094"
}

var addCmd = &cli.Command{
	Name:  "add",
	Usage: "Add the peer address to the bootstrap persistent",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "api",
			Aliases: []string{"a"},
			Usage:   "Endpoint of the cluster restful API",
			Value:   defaultClusterAPIEndpoint(),
		},
		&cli.StringFlag{
			Name:    "dir",
			Aliases: []string{"d"},
			Usage:   "Directory to the IPFS repo",
			Value:   defaultRepoPath(),
		},
	},
	Action: func(c *cli.Context) error {
		return errors.New("not available now")
	},
}
