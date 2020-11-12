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
		addrCmd,
		exitCmd,
		helpCmd,
		addCmd,
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

var addrCmd = &cli.Command{
	Name:  "addr",
	Usage: "Show the address of current drive",
	Action: func(c *cli.Context) error {
		d, err := getDrive(c)
		if err != nil {
			return err
		}

		fmt.Println(d.Address())
		return nil
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
	Name:      "ls",
	Aliases:   []string{"l"},
	Usage:     "List all existing file on this drive",
	UsageText: "ls <prefix>",
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

		if len(lr.Files()) == 0 {
			fmt.Println("empty result")
			return nil
		}

		_, err = lr.WriteTo(os.Stdout)
		return err
	},
}

var addCmd = &cli.Command{
	Name:      "add",
	Aliases:   []string{"a"},
	Usage:     "Add local file to current drive",
	UsageText: "add <key> <path>",
	Before: func(c *cli.Context) error {
		return nil
	},
	Action: func(c *cli.Context) error {
		var (
			ctx   = c.Context()
			key   = c.Args()[0]
			fpath = c.Args()[1]
		)

		d, err := getDrive(c)
		if err != nil {
			return err
		}

		f, err := d.Add(ctx, key, fpath)
		if err != nil {
			return err
		}

		fmt.Printf("Add file %s (%s) as %s\n", f.Key, f.Cid, sizeToStr(f.Size))

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
