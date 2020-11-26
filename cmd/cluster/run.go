package cluster

import (
	"errors"
	"strings"

	cls "github.com/meowdada/ipfstor/cluster"
	"github.com/urfave/cli/v2"
)

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "Run the cluster peer with given settings",
	Flags: []cli.Flag{
		configFlag,
		identityFlag,
		&cli.BoolFlag{
			Name:    "detach",
			Aliases: []string{"d"},
			Usage:   "Running this cluster in detached mode",
		},
		&cli.StringFlag{
			Name:  "peers",
			Usage: "Address of peers to be bootstrap. Seperated by comma.",
		},
	},
	Action: func(c *cli.Context) error {
		ctx := c.Context
		detached := c.Bool("detach")
		configPath, identityPath := getConfigs(c)
		peerStr := c.String("peers")
		peers := strings.Split(peerStr, ",")
		if len(peerStr) == 0 {
			peers = nil
		}

		cluster, err := cls.New(ctx, configPath, identityPath, peers)
		if err != nil {
			return err
		}

		if err := cluster.Bootstrap(ctx, peers); err != nil {
			return err
		}

		if detached {
			return errors.New("not yet support detached mode")
		}

		return cluster.Run(ctx, peers)
	},
}
