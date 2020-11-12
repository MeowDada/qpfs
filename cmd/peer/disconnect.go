package peer

import (
	"fmt"

	ma "github.com/multiformats/go-multiaddr"
	"github.com/urfave/cli/v2"
)

var disconnectCmd = &cli.Command{
	Name:    "disconnect",
	Aliases: []string{"d"},
	Usage:   "Disconnect to a connected peer",
	Flags: []cli.Flag{
		apiFlag,
	},
	Action: func(c *cli.Context) error {
		var (
			ctx   = c.Context
			addrs = c.Args().Slice()
		)

		api, err := getAPI(c)
		if err != nil {
			return err
		}

		for _, addr := range addrs {
			maAddr, err := ma.NewMultiaddr(addr)
			if err != nil {
				fmt.Println(err)
				continue
			}

			err = api.Swarm().Disconnect(ctx, maAddr)
			if err != nil {
				return err
			}
			fmt.Printf("disconnect to %s\n", maAddr.String())
		}

		return nil
	},
}
