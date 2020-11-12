package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var idCmd = &cli.Command{
	Name:  "id",
	Usage: "Show current identity of the ipfs node",
	Flags: []cli.Flag{
		apiFlag,
	},
	Action: func(c *cli.Context) error {
		ctx := c.Context

		d, err := getAPI(c)
		if err != nil {
			return err
		}

		k, err := d.Key().Self(ctx)
		if err != nil {
			return err
		}

		fmt.Println(k.ID())
		return err
	},
}
