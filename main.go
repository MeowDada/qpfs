package main

import (
	"github.com/meowdada/qpfs/cmd"

	// Import subcommands.
	_ "github.com/meowdada/qpfs/cmd/drive"
)

var Version string

func main() {
	cmd.Main(Version)
}
