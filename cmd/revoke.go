package cmd

import (
	"fmt"

	"github.com/meowdada/ipfstor/drive"
	"github.com/meowdada/ipfstor/options"
	"github.com/urfave/cli/v2"
)

var revokeCmd = &cli.Command{
	Name:      "revoke",
	Aliases:   []string{"r"},
	Usage:     "Revoke write permission from a user",
	UsageText: "qpfs revoke <resolve> <userID>",
	Flags: []cli.Flag{
		apiFlag,
		dirFlag,
	},
	Before: func(c *cli.Context) error {
		if c.Args().Len() != 2 {
			return fmt.Errorf("usage: qpfs revoke <resolve> <userID>")
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

		if err := d.Revoke(ctx, userID, "write"); err != nil {
			return err
		}

		fmt.Println("revoke write permission from user", userID)

		return nil
	},
}
