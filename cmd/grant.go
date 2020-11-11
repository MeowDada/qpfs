package cmd

import (
	"fmt"

	"github.com/meowdada/ipfstor/drive"
	"github.com/meowdada/ipfstor/options"
	"github.com/urfave/cli/v2"
)

var grantCmd = &cli.Command{
	Name:      "grant",
	Aliases:   []string{"g"},
	Usage:     "Grant write permission to specific user",
	UsageText: "qpfs grant <resolve> <userID>",
	Flags: []cli.Flag{
		apiFlag,
		dirFlag,
	},
	Before: func(c *cli.Context) error {
		if c.Args().Len() != 2 {
			return fmt.Errorf("usage: qpfs grant <user> <permission>")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		var (
			ctx     = c.Context
			dir     = c.String("dir")
			args    = c.Args().Slice()
			resolve = args[0]
			userID  = args[1]
		)

		api, err := getAPI(c)
		if err != nil {
			return err
		}

		opts := options.OpenDrive().
			SetDirectory(dir)

		d, err := drive.Open(ctx, api, resolve, opts)
		if err != nil {
			return err
		}
		defer d.Close(ctx)

		if err := d.Grant(ctx, userID, "write"); err != nil {
			return err
		}

		fmt.Println("grant write permission to user", userID)

		return nil
	},
}
