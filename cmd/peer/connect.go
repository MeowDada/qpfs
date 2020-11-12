package peer

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var connectCmd = &cli.Command{
	Name:    "connect",
	Aliases: []string{"c"},
	Usage:   "Connect to an existing peer",
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

		infos, err := parseAddresses(ctx, addrs)
		if err != nil {
			return err
		}

		for _, info := range infos {
			err := api.Swarm().Connect(ctx, info)
			if err != nil {
				return err
			}
			fmt.Printf("connect to %s successfully\n", info.String())
		}

		return nil
	},
}
