package shell

import (
	"errors"
	"fmt"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/meowdada/ipfstor/drive"
	"github.com/meowdada/qpfs/pkg/cli"
)

// App is a pre-defined shell instance.
var App = &cli.Shell{
	Commands: []*cli.Command{
		lsCmd,
		exitCmd,
		helpCmd,
	},
}

// ErrExitShell is a special error that hints to exit the cli progarm.
var ErrExitShell = errors.New("exit shell error")

var helpCmd = &cli.Command{
	Name:    "help",
	Aliases: []string{"h"},
	Usage:   "Show usage help text",
	Action: func(c *cli.Context) error {
		return c.Shell().WriteUsage(os.Stdout)
	},
}

var exitCmd = &cli.Command{
	Name:  "exit",
	Usage: "Exit this interactive cli program",
	Action: func(c *cli.Context) error {
		return ErrExitShell
	},
}

var lsCmd = &cli.Command{
	Name:    "ls",
	Aliases: []string{"l"},
	Usage:   "List all existing file on this drive",
	Before: func(c *cli.Context) error {
		if len(c.Args()) > 1 {
			return fmt.Errorf("usage: ls <prefix>")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		prefix := ""
		if len(c.Args()) > 0 {
			prefix = c.Args()[0]
		}

		d, err := getDrive(c)
		if err != nil {
			return err
		}

		lr, err := d.List(c.Context(), prefix)
		if err != nil {
			return err
		}

		for _, f := range lr.Files() {
			fmt.Printf("%s (%s): %s\n", f.Key, f.Cid, sizeToStr(f.Size))
		}

		return nil
	},
}

func sizeToStr(size int64) string {
	return humanize.IBytes(uint64(size))
}

func getDrive(c *cli.Context) (drive.Instance, error) {
	d, ok := c.Get("drive")
	if !ok {
		return nil, fmt.Errorf("no available drive instance")
	}
	return d.(drive.Instance), nil
}
