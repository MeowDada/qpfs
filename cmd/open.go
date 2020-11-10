package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/meowdada/ipfstor/drive"
	"github.com/meowdada/ipfstor/ipfsutil"
	"github.com/meowdada/ipfstor/options"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
)

func defaultOrbitDBPath() string {
	dir, _ := homedir.Dir()
	return filepath.Join(dir, ".orbitdb")
}

var openCmd = &cli.Command{
	Name:  "open",
	Usage: "Open an existing drive by given name of the database or its address",
	Before: func(c *cli.Context) error {
		if c.Args().Len() != 1 {
			return fmt.Errorf("usage: qpfs open <resolve>")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		var (
			ctx     = c.Context
			resolve = c.Args().First()
		)

		api, err := ipfsutil.NewAPI(ipfsutil.DefaultAPIAddress)
		if err != nil {
			return err
		}

		d, err := drive.Open(ctx, api, resolve, options.OpenDrive().SetDirectory(defaultOrbitDBPath()))
		if err != nil {
			return err
		}
		defer d.Close(ctx)

		fmt.Printf("Open drive: %s\n", d.Address())

		for {
		}

		return nil
	},
}
