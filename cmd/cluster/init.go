package cluster

import (
	"fmt"

	cls "github.com/meowdada/ipfstor/cluster"
	"github.com/meowdada/ipfstor/options"
	"github.com/urfave/cli/v2"
)

var initCmd = &cli.Command{
	Name:  "init",
	Usage: "Initialize a new cluster config",
	Flags: []cli.Flag{
		configFlag,
		identityFlag,
		&cli.BoolFlag{
			Name:    "random-ports",
			Aliases: []string{"r"},
			Usage:   "Using random ports to host the cluster related services",
			Value:   false,
		},
		&cli.StringFlag{
			Name:    "secret",
			Aliases: []string{"s"},
			Usage:   "Sets a secret key of the cluster",
		},
	},
	Action: func(c *cli.Context) error {
		ctx := c.Context
		configPath, identityPath := getConfigs(c)
		useRandomPorts := c.Bool("random-ports")
		secret := c.String("secret")

		if err := cls.NewConfig(ctx, options.Cluster().
			SetConfigPath(configPath).
			SetIdentityPath(identityPath).
			SetRandomPorts(useRandomPorts).
			SetSecret(secret),
		); err != nil {
			return err
		}

		fmt.Println("Saving config file as", configPath)
		fmt.Println("Saving identity file as", identityPath)

		return nil
	},
}
