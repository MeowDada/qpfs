package main

import (
	"github.com/meowdada/qpfs/cmd"

	// Import subcommands.
	_ "github.com/meowdada/qpfs/cmd/drive"
	_ "github.com/meowdada/qpfs/cmd/peer"
	_ "github.com/meowdada/qpfs/cmd/util"
)

var (
	version string
)

func main() {
	cmd.Main(version)
}
