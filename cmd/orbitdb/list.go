package orbitdb

import (
	"fmt"

	"github.com/meowdada/qpfs/drive"
	"github.com/urfave/cli/v2"
)

var listCmd = &cli.Command{
	Name:  "ls",
	Usage: "List all existing file in this drive",
	Action: func(c *cli.Context) error {
		api, err := drive.NewClient("http://localhost:9090")
		if err != nil {
			return err
		}

		lr, err := api.List(c.Context, "")
		if err != nil {
			return err
		}

		fmt.Println(lr)
		return nil
	},
}
