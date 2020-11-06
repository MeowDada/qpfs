package drive

import (
	"fmt"

	"github.com/meowdada/ipfstor"
	"github.com/meowdada/qpfs/pkg/ipfs"
	"github.com/urfave/cli/v2"
)

var openCmd = &cli.Command{
	Name:      "open",
	Usage:     "Opens an existing drive on the IPFS netowrk.",
	UsageText: "qpfs open <address>",
	Description: `Opens an existing drive on the IPFS network. For example:

	qpfs open /orbitdb/bafyreigiemkcw7ildl5pnblp2hc7g7m3d5pwe6sw3bpodzdew2f4jybptm/123

It will open a drive with human readable name 123.
`,
	Before: func(c *cli.Context) error {
		if c.Args().Len() != 1 {
			return fmt.Errorf("expect only 1 address")
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

		d, err := ipfstor.OpenDrive(ctx, api, addr)
		if err != nil {
			return err
		}

		return d.Close(ctx)
	},
}
