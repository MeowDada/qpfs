package cmd

import (
	"fmt"

	"github.com/meowdada/ipfstor"
	"github.com/meowdada/qpfs/pkg/ipfs"
	cli "github.com/urfave/cli/v2"
)

var rmCmd = &cli.Command{
	Name:      "rm",
	Usage:     "Unpin specific file on the drive",
	UsageText: "qpfs rm <drive>/<file>",
	Description: `Unpin specific file on the drive. For example:
	
	qpfs rm abc/123.txt

This will unpin a file called 123.txt on the drive abc if it exists.`,
	Before: func(c *cli.Context) error {
		if c.Args().Len() != 1 {
			return fmt.Errorf("usage:  qpfs rm <drive>/<file> to remove the file")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		var (
			ctx  = c.Context
			addr = "/ip4/127.0.0.1/tcp/5001"
		)

		if c.String("api") != "" {
			addr = c.String("api")
		}

		api, err := ipfs.NewAPI(addr)
		if err != nil {
			return err
		}

		drive, obj, err := ipfs.ParseDriveObject(c.Args().First())
		if err != nil {
			return err
		}

		d, err := ipfstor.OpenDrive(ctx, api, drive)
		if err != nil {
			return err
		}
		defer d.Close(ctx)

		return d.Remove(ctx, obj)
	},
}
