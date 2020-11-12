package peer

import (
	"fmt"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var lsCmd = &cli.Command{
	Name:    "ls",
	Aliases: []string{"l"},
	Usage:   "List all connected address",
	Flags: []cli.Flag{
		apiFlag,
	},
	Action: func(c *cli.Context) error {
		ctx := c.Context

		api, err := getAPI(c)
		if err != nil {
			return err
		}

		infos, err := api.Swarm().Peers(ctx)
		if err != nil {
			return err
		}

		for i := range infos {
			id := infos[i].ID()
			addr := infos[i].Address()
			fmt.Println(filepath.Join(addr.String(), "p2p", id.String()))
		}
		return nil
	},
}
