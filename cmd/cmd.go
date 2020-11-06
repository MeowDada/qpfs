package cmd

import (
	"log"
	"os"
	"sort"

	cli "github.com/urfave/cli/v2"
)

const (
	catagoryFileOps = "file operations"
)

var mainApp = &cli.App{
	Name:  "qpfs",
	Usage: "A cli tool to interacts with ipfs share drives.",
	Description: `A cli tool to interacts with ipfs share drives.

This application is based on private IPFS network, which means it
is isolated from the public IPFS network. This feature make routing
faster and safer to shares files between trust nodes.

This cli requires at least a running IPFS node to work. And it also
expected the default repo path located at $HOME/.ipfs. If you use custom
repo settings, you might need to configure the below environment variable:

	IPFS_PATH=<path-to-repo>


	`,
	Commands: []*cli.Command{
		addCmd,
		getCmd,
		lsCmd,
		rmCmd,
	},
}

// Register register commands to main app.
func Register(cmds ...*cli.Command) {
	// Sort the command in alphabet order.
	sort.Slice(cmds, func(i, j int) bool {
		return cmds[i].Name < cmds[j].Name
	})
	mainApp.Commands = append(mainApp.Commands, cmds...)
}

// Main is the entrypoint of the cli.
func Main(ver string) {
	// Sort the command in alphabet order.
	sort.Slice(mainApp.Commands, func(i, j int) bool {
		return mainApp.Commands[i].Name < mainApp.Commands[j].Name
	})

	mainApp.Version = ver

	if err := mainApp.Run(os.Args); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
