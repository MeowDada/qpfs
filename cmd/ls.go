package cmd

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/meowdada/ipfstor/drive"
	"github.com/meowdada/ipfstor/options"
	"github.com/urfave/cli/v2"
)

var lsCmd = &cli.Command{
	Name:      "ls",
	Aliases:   []string{"l"},
	Usage:     "List all existing file in the drive which matches given prefix",
	UsageText: "qpfs ls <prefix>",
	Flags: []cli.Flag{
		apiFlag,
		dirFlag,
	},
	Before: func(c *cli.Context) error {
		if c.Args().Len() > 2 || c.Args().Len() == 0 {
			return fmt.Errorf("usage: qpfs ls <resolve> <prefix>")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		var (
			ctx     = c.Context
			dir     = c.String("dir")
			resolve = c.Args().First()
			prefix  string
		)

		if c.Args().Len() > 1 {
			prefix = c.Args().Slice()[1]
		}

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

		lr, err := d.List(ctx, prefix)
		if err != nil {
			return err
		}

		fs := lr.Files()

		for i := range fs {
			fmt.Printf("%s (%s): %s\n", fs[i].Key, fs[i].Cid, humanize.IBytes(uint64(fs[i].Size)))
		}

		return nil
	},
}
