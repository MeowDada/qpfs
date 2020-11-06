package drive

import (
	"context"
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/ipfs/go-cid"
	"github.com/meowdada/ipfstor"
	"github.com/meowdada/qpfs/pkg/ipfs"
	"github.com/urfave/cli/v2"
)

var statCmd = &cli.Command{
	Name:        "stat",
	Usage:       "Print out the information about specific drive",
	Description: "",
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
			d, err := ipfstor.OpenDrive(ctx, api, drive)
			if err != nil {
				fmt.Printf("unable to open drive %s: %v\n", drive, err)
				continue
			} else {
				var count int
				var sum int64

				d.Iter(ctx, func(ctx context.Context, key string, cid cid.Cid, size int64, owner string) error {
					sum += size
					count++
					return nil
				})
				fmt.Printf("Address: %s\n", d.Address())
				fmt.Printf("Size:    %s (%d bytes)\n", humanize.IBytes(uint64(sum)), sum)
				fmt.Printf("#Files:  %d\n", count)
				d.Close(ctx)
			}
		}

		return nil
	},
}
