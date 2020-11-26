package cluster

import "github.com/urfave/cli/v2"

var joinCmd = &cli.Command{
	Name:  "join",
	Usage: "Join an existing cluster peer",
	Flags: []cli.Flag{
		configFlag,
		identityFlag,
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
