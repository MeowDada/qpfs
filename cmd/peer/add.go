package peer

import (
	"fmt"
	"path/filepath"

	ipfscluster "github.com/ipfs/ipfs-cluster"
	"github.com/meowdada/ipfstor/cluster"
	"github.com/mitchellh/go-homedir"
	ma "github.com/multiformats/go-multiaddr"
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
		var (
			ctx  = c.Context
			args = c.Args().Slice()
		)

		cls, err := cluster.New(ctx, "/ip4/127.0.0.1/tcp/9094")
		if err != nil {
			return err
		}

		for _, p := range args {
			addr, err := ma.NewMultiaddr(p)
			if err != nil {
				fmt.Printf("Parse multiaddr from %s with error: %v\n", p, err)
				continue
			}

			ids := ipfscluster.PeersFromMultiaddrs([]ma.Multiaddr{addr})

			_, err = cls.PeerAdd(ctx, ids[0])
			if err != nil {
				fmt.Printf("Add peer %v with error: %v\n", p, err)
			}
		}

		return nil
	},
}
