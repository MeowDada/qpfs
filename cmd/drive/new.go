package drive

import (
	"fmt"

	"github.com/meowdada/ipfstor"
	"github.com/meowdada/qpfs/pkg/ipfs"
	"github.com/urfave/cli/v2"
)

var newCmd = &cli.Command{
	Name:      "new",
	Usage:     "Create an new drive on local",
	UsageText: "qpfs drive new [drive...]",
	Description: `Follwoing example create a new drive named drive1:

	qpfs drive new dirve1

It is also possible to create multiple drives at once as below example:

	qpfs drive new drive1 drive2 drive3

There are totally 3 drives will be created.`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "api",
			Aliases: []string{"a"},
			Usage:   "Specific the `address` to the api which is going to be used as backend",
		},
	},
	Before: func(c *cli.Context) error {
		if c.Args().Len() < 1 {
			return fmt.Errorf("must specific at least 1 drive")
		}
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

		drives := c.Args().Slice()
		for i := range drives {
			drive := drives[i]
			d, err := ipfstor.NewDriver(ctx, api, drive)
			if err != nil {
				fmt.Printf("create drive %s with error: %v\n", drive, err)
			} else {
				fmt.Printf("create drive %s successfully\n", drive)
			}
			d.Close(ctx)
		}
		return nil
	},
}
