package cmd

import (
	"fmt"

	"github.com/meowdada/ipfstor"
	"github.com/meowdada/qpfs/pkg/ipfs"
	cli "github.com/urfave/cli/v2"
)

var addCmd = &cli.Command{
	Name:      "add",
	Usage:     "Upload file to the specific drive",
	UsageText: "qpfs add <file> <drive>",
	Description: `Upload file to the specific drive. To any upload files,
Only basename will be preserved.`,
	Before: func(c *cli.Context) error {
		if c.Args().Len() != 2 {
			return fmt.Errorf("usage: qpfs add <src> <dst>")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		var (
			ctx   = c.Context
			addr  = "/ip4/127.0.0.1/tcp/5001"
			args  = c.Args().Slice()
			fpath = args[0]
		)

		if c.String("api") != "" {
			addr = c.String("api")
		}

		api, err := ipfs.NewAPI(addr)
		if err != nil {
			return err
		}

		path, err := ipfs.NewDstPath(fpath, args[1])
		if err != nil {
			return err
		}

		d, err := ipfstor.OpenDrive(ctx, api, path.Drive())
		if err != nil {
			return err
		}
		defer d.Close(ctx)

		return d.Add(ctx, path.Filename(), fpath)
	},
}
