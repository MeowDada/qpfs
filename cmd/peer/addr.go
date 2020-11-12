package peer

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var addrCmd = &cli.Command{
	Name:  "addr",
	Usage: "Print out current address informations",
	Flags: []cli.Flag{
		apiFlag,
	},
	Action: func(c *cli.Context) error {
		var (
			ctx = c.Context
		)

		api, err := getAPI(c)
		if err != nil {
			return err
		}

		addrs, err := api.Swarm().LocalAddrs(ctx)
		if err != nil {
			return err
		}

		for i := range addrs {
			fmt.Println(addrs[i].String())
		}

		return nil
	},
}
