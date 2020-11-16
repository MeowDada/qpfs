package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/meowdada/ipfstor/drive"
	"github.com/meowdada/ipfstor/options"
	"github.com/urfave/cli/v2"
)

var addCmd = &cli.Command{
	Name:      "add",
	Aliases:   []string{"a"},
	Usage:     "Add local file to the specific drive",
	UsageText: "qpfs add <resolve> <file>",
	Flags: []cli.Flag{
		apiFlag,
		dirFlag,
	},
	Before: func(c *cli.Context) error {
		if c.Args().Len() != 2 {
			return fmt.Errorf("usage: qpfs add <resolve> <file>")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		var (
			ctx     = c.Context
			dir     = c.String("dir")
			args    = c.Args().Slice()
			fpath   = args[1]
			resolve = args[0]
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

		key := filepath.Base(fpath)
		info, err := d.AddFile(ctx, key, fpath)
		if err != nil {
			return err
		}

		fmt.Printf("Add %s %s\n", key, info.Cid)
		return nil
	},
}
