package peer

import (
	"fmt"
	"path/filepath"

	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/meowdada/ipfstor/cluster"
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
		var (
			ctx  = c.Context
			args = c.Args().Slice()
		)

		cls, err := cluster.New(ctx, "/ip4/127.0.0.1/tcp/9094")
		if err != nil {
			return err
		}

		ids, err := cls.Peers(ctx)
		if err != nil {
			return err
		}

		for _, id := range ids {
			fmt.Println(id.ID)
		}

		for _, p := range args {
			id, err := peer.Decode(p)
			if err != nil {
				fmt.Printf("Parse peer ID from string %s with error: %v\n", p, err)
				return err
			}

			_, err = cls.PeerAdd(ctx, id)
			if err != nil {
				fmt.Printf("Add peer %v with error: %v\n", id.String(), err)
			}
		}

		return nil
	},
}
