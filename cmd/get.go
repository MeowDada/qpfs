package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/meowdada/ipfstor/drive"
	"github.com/meowdada/ipfstor/options"
	"github.com/urfave/cli/v2"
)

var getCmd = &cli.Command{
	Name:      "get",
	Usage:     "Download file from current drive to specific path",
	UsageText: "qpfs get <resolve> <key>",
	Flags: []cli.Flag{
		dirFlag,
		apiFlag,
		&cli.StringFlag{
			Name:    "path",
			Aliases: []string{"o"},
			Usage:   "Path to the download file",
		},
	},
	Before: func(c *cli.Context) error {
		if c.Args().Len() != 2 {
			return fmt.Errorf("usage: qpfs get <resolve> <key>")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		var (
			ctx     = c.Context
			args    = c.Args().Slice()
			resolve = args[0]
			key     = args[1]
			dir     = c.String("dir")
			path    = c.String("path")
		)

		if len(path) == 0 {
			path = filepath.Base(key)
		}

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

		f, err := d.Get(ctx, key)
		if err != nil {
			return err
		}
		defer f.Close()

		dst, err := os.Create(path)
		if err != nil {
			return err
		}
		defer dst.Close()

		_, err = io.Copy(dst, f)
		if err != nil {
			return err
		}

		fmt.Printf("Download %s to %s\n", key, path)

		return nil
	},
}
