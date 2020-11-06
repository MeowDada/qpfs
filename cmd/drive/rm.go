package drive

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var rmCmd = &cli.Command{
	Name:  "rm",
	Usage: "Remove local drives",
	Description: `rm will remove local drives. It will not remove a local drive
if it is not empty`,
	Action: func(c *cli.Context) error {
		return fmt.Errorf("not implement yet")
	},
}
