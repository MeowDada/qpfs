package orbitdb

import (
	"fmt"

	"github.com/meowdada/qpfs/drive"
	"github.com/urfave/cli/v2"
)

var addCmd = &cli.Command{
	Name:  "add",
	Usage: "Add file to ipfs drive",
	Before: func(c *cli.Context) error {
		if c.Args().Len() != 2 {
			return fmt.Errorf("usage: qpfs orbitdb add <src> <dst>")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		var (
			src = c.Args().Slice()[0]
			dst = c.Args().Slice()[1]
		)

		api, err := drive.NewClient("http://localhost:9090")
		if err != nil {
			return err
		}

		return api.Add(c.Context, dst, src)
	},
}
