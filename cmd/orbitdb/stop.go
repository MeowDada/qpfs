package orbitdb

import (
	"github.com/meowdada/qpfs/drive"
	"github.com/urfave/cli/v2"
)

var stopCmd = &cli.Command{
	Name:  "stop",
	Usage: "Stop the corresponding ipfs daemon",
	Action: func(c *cli.Context) error {
		api, err := drive.NewClient("http://localhost:9090")
		if err != nil {
			return err
		}

		return api.Stop(c.Context)
	},
}
