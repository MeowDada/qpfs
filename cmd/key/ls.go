package key

import (
	"io"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
)

const swarmKeyFname = "swarm.key"

func defaultIPFSRepoPath() string {
	dir, _ := homedir.Dir()
	return filepath.Join(dir, ".ipfs")
}

var lsCmd = &cli.Command{
	Name:  "ls",
	Usage: "Show current swarm key if it present",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "dir",
			Usage: "Path to the ipfs node repo",
			Value: defaultIPFSRepoPath(),
		},
	},
	Action: func(c *cli.Context) error {
		var (
			dir = c.String("dir")
		)

		path := filepath.Join(dir, swarmKeyFname)

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(os.Stdout, f)
		return err
	},
}
