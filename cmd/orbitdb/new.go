package orbitdb

import (
	"os"
	"path/filepath"

	"github.com/meowdada/ipfstor"
	"github.com/meowdada/qpfs/drive"
	"github.com/meowdada/qpfs/pkg/ipfs"
	"github.com/urfave/cli/v2"
)

var (
	defaultDBPath = filepath.Join(os.Getenv("HOME"), ".orbitdb")
)

var newCmd = &cli.Command{
	Name:        "new",
	Usage:       "Create a new orbitdb",
	UsageText:   "qpfs orbitdb new <database>",
	Description: "",
	Before: func(c *cli.Context) error {
		return nil
	},
	Action: func(c *cli.Context) error {
		ctx := c.Context
		addr := "/ip4/127.0.0.1/tcp/5001"

		if c.String("api") != "" {
			addr = c.String("api")
		}

		api, err := ipfs.NewAPI(addr)
		if err != nil {
			return err
		}

		driver, err := ipfstor.NewDriver(ctx, api, c.Args().First())
		if err != nil {
			return err
		}

		if err := driver.Close(ctx); err != nil {
			return err
		}

		daemon, err := drive.NewDaemon(":9090", driver.Address())
		if err != nil {
			return err
		}

		return daemon.Start()
	},
}

func boolPtr(flag bool) *bool { return &flag }

func strPtr(str string) *string { return &str }
