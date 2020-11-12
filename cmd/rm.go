package cmd

import (
	"fmt"

	"github.com/meowdada/ipfstor/drive"
	"github.com/meowdada/ipfstor/options"
	"github.com/urfave/cli/v2"
)

var rmCmd = &cli.Command{
	Name:      "rm",
	Usage:     "Remove file on spcific drive",
	UsageText: "qpfs rm <drive> <key>",
	Flags: []cli.Flag{
		dirFlag,
		apiFlag,
	},
	Before: func(c *cli.Context) error {
		if c.Args().Len() != 2 {
			return fmt.Errorf("usage: qpfs rm <drive> <key>")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		var (
			ctx     = c.Context
			dir     = c.String("dir")
			args    = c.Args().Slice()
			resolve = args[0]
			key     = args[1]
		)

		api, err := getAPI(c)
		if err != nil {
			return err
		}

		opts := options.OpenDrive().
			SetDirectory(dir).
			SetCreate(false)

		d, err := drive.Open(ctx, api, resolve, opts)
		if err != nil {
			return err
		}
		defer d.Close(ctx)

		return d.Remove(ctx, key)
	},
}
