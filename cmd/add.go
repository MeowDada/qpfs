package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/meowdada/ipfstor/drive"
	"github.com/meowdada/ipfstor/ipfsutil"
	"github.com/meowdada/ipfstor/options"
	"github.com/urfave/cli/v2"
)

var addCmd = &cli.Command{
	Name:  "add",
	Usage: "Add local file to the specific drive",
	Before: func(c *cli.Context) error {
		if c.Args().Len() != 2 {
			return fmt.Errorf("usage: qpfs add <file> <resolve>")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		var (
			ctx     = c.Context
			args    = c.Args().Slice()
			fpath   = args[0]
			resolve = args[1]
		)

		api, err := ipfsutil.NewAPI(ipfsutil.DefaultAPIAddress)
		if err != nil {
			return err
		}

		opts := options.OpenDrive().
			SetDirectory(defaultOrbitDBPath())

		d, err := drive.Open(ctx, api, resolve, opts)
		if err != nil {
			return err
		}
		defer d.Close(ctx)

		key := filepath.Base(fpath)
		info, err := d.Add(ctx, key, fpath)
		if err != nil {
			return err
		}

		fmt.Printf("Add %s %s\n", key, info.Cid)
		return nil
	},
}
