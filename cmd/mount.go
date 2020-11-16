package cmd

import (
	"fmt"

	"github.com/meowdada/ipfstor/drive"
	"github.com/meowdada/ipfstor/fs"
	"github.com/meowdada/ipfstor/options"
	"github.com/urfave/cli/v2"
)

var mountCmd = &cli.Command{
	Name:      "mount",
	Aliases:   []string{"m"},
	Usage:     "Mount a drive to specific path",
	UsageText: "qpfs mount <drive> <dir>",
	Flags: []cli.Flag{
		apiFlag,
		dirFlag,
	},
	Before: func(c *cli.Context) error {
		if c.Args().Len() != 2 {
			return fmt.Errorf("usage: qpfs mount <drive> <dir>")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		var (
			ctx        = c.Context
			dir        = c.String("dir")
			args       = c.Args().Slice()
			drivePath  = args[0]
			mountpoint = args[1]
		)

		api, err := getAPI(c)
		if err != nil {
			return err
		}

		opts := options.OpenDrive().
			SetDirectory(dir).
			SetCreate(false)

		d, err := drive.Open(ctx, api, drivePath, opts)
		if err != nil {
			return err
		}
		defer d.Close(ctx)

		return fs.Mount(mountpoint, d)
	},
}
