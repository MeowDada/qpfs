package cmd

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/meowdada/ipfstor"
	"github.com/meowdada/qpfs/pkg/ipfs"
	"github.com/urfave/cli/v2"
)

var lsCmd = &cli.Command{
	Name:        "ls",
	Usage:       "List all existing files under the drive",
	UsageText:   "qpfs ls <drive>",
	Description: "List all existing files under the drive. If the driver contains no files, nothing will be returned",
	Before: func(c *cli.Context) error {
		if c.Args().Len() != 1 {
			return fmt.Errorf("specific one drive to list files")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		var (
			ctx   = c.Context
			addr  = "/ip4/127.0.0.1/tcp/5001"
			drive = c.Args().First()
		)

		if c.String("api") != "" {
			addr = c.String("api")
		}

		api, err := ipfs.NewAPI(addr)
		if err != nil {
			return err
		}

		d, err := ipfstor.OpenDrive(ctx, api, drive)
		if err != nil {
			return err
		}
		defer d.Close(ctx)

		rs, err := d.List(ctx)
		if err != nil {
			return err
		}

		for i := range rs {
			fmt.Printf("%s (%s)\n", rs[i].Key, humanize.IBytes(uint64(rs[i].Size)))
		}

		return nil
	},
}
