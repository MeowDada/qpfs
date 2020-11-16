package shell

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/meowdada/ipfstor/drive"
	"github.com/meowdada/qpfs/pkg/cli"
)

// App is a pre-defined shell instance.
var App = &cli.Shell{
	Commands: []*cli.Command{
		lsCmd,
		rmCmd,
		addrCmd,
		addCmd,
		getCmd,
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

		b := lr.Bytes(drive.ListMaskKey | drive.ListMaskSize | drive.ListMaskTime)
		fmt.Println(string(b))
		return nil
	},
}

var rmCmd = &cli.Command{
	Name:      "rm",
	Aliases:   []string{"r"},
	Usage:     "Remove file on the current drive",
	UsageText: "rm <key>",
	Before: func(c *cli.Context) error {
		return nil
	},
	Action: func(c *cli.Context) error {
		var (
			ctx  = c.Context()
			args = c.Args()
		)

		d, err := getDrive(c)
		if err != nil {
			return err
		}

		for i := range args {
			if err := d.Remove(ctx, args[i]); err != nil {
				fmt.Printf("remove %s with error: %v\n", args[i], err)
				continue
			}
			fmt.Printf("remove %s successfully\n", args[i])
		}

		return nil
	},
}

var addCmd = &cli.Command{
	Name:      "add",
	Aliases:   []string{"a"},
	Usage:     "Add local file to current drive",
	UsageText: "add <key> <path>",
	Before: func(c *cli.Context) error {
		if len(c.Args()) != 2 {
			return fmt.Errorf("usage: add <key> <path>")
		}
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

		f, err := d.AddFile(ctx, key, fpath)
		if err != nil {
			return err
		}

		fmt.Printf("Add file %s (%s) as %s\n", f.Key, f.Cid, sizeToStr(f.Size))

		return nil
	},
}

var getCmd = &cli.Command{
	Name:      "get",
	Aliases:   []string{"g"},
	Usage:     "Download file from the current drive",
	UsageText: "get <key> <path>",
	Before: func(c *cli.Context) error {
		if len(c.Args()) != 2 {
			return fmt.Errorf("usage: get <key> <path>")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		var (
			ctx  = c.Context()
			args = c.Args()
			key  = args[0]
			path = args[1]
		)

		d, err := getDrive(c)
		if err != nil {
			return err
		}

		r, err := d.Get(ctx, key)
		if err != nil {
			return err
		}

		f, err := os.Create(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, r)
		if err != nil {
			return err
		}

		fmt.Printf("Download %s to %s\n", key, path)

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
