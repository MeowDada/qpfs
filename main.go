package main

import (
	"github.com/meowdada/qpfs/cmd"

	// Import drive plugin commands.
	_ "github.com/meowdada/qpfs/cmd/drive"
)

var Version string

func main() {
	cmd.Main(Version)
}
