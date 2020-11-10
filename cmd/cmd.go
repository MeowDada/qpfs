package cmd

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

var mainApp = &cli.App{
	Name:  "qpfs",
	Usage: "A cli tool to interact with ipfs drive",
	Commands: []*cli.Command{
		openCmd,
		addCmd,
		lsCmd,
	},
}

// Main is the entrypoint
func Main(ver string) {
	mainApp.Version = ver
	if err := mainApp.Run(os.Args); err != nil {
		log.Println(err)
		return
	}
}

// Register registers multiple commands to main app.
func Register(cmds ...*cli.Command) {
	sort.Slice(cmds, func(i, j int) bool {
		return cmds[i].Name < cmds[j].Name
	})
	mainApp.Commands = append(mainApp.Commands, cmds...)
}
