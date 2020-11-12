package main

import (
	"github.com/meowdada/qpfs/cmd"

	// Import subcommands.
	_ "github.com/meowdada/qpfs/cmd/drive"
	_ "github.com/meowdada/qpfs/cmd/peer"
)

var Version string

func main() {
	cmd.Main(Version)
}
