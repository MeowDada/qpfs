package cluster

import (
	cls "github.com/meowdada/ipfstor/cluster"
	"github.com/meowdada/qpfs/cmd"
	"github.com/urfave/cli/v2"
)

func init() {
	cmd.Register(clusterCmd)
}

var clusterCmd = &cli.Command{
	Name:  "cluster",
	Usage: "IPFS Cluster management",
	Subcommands: []*cli.Command{
		initCmd,
		runCmd,
		joinCmd,
	},
}

var configFlag = &cli.StringFlag{
	Name:    "config",
	Aliases: []string{"c"},
	Usage:   "Path to the ipfs cluster config file",
	Value:   cls.DefaultConfigPath(),
}

var identityFlag = &cli.StringFlag{
	Name:    "id",
	Aliases: []string{"i"},
	Usage:   "Path to the identity file",
	Value:   cls.DefaultIdentityPath(),
}

func getConfigs(c *cli.Context) (string, string) {
	return c.String("config"), c.String("id")
}
