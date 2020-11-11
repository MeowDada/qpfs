package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/meowdada/ipfstor/drive"
	"github.com/meowdada/ipfstor/options"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
)

func defaultOrbitDBPath() string {
	dir, _ := homedir.Dir()
	return filepath.Join(dir, ".orbitdb")
}

var openCmd = &cli.Command{
	Name:      "open",
	Aliases:   []string{"o"},
	Usage:     "Open an existing drive by given name of the database or its address",
	UsageText: "qpfs open <resolve>",
	Flags: []cli.Flag{
		apiFlag,
		dirFlag,
		&cli.BoolFlag{
			Name:    "memory",
			Aliases: []string{"m"},
			Usage:   "Using memory mode to host the datastore",
		},
	},
	Before: func(c *cli.Context) error {
		if c.Args().Len() != 1 {
			return fmt.Errorf("usage: qpfs open <resolve>")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		var (
			ctx        = c.Context
			resolve    = c.Args().First()
			dir        = c.String("dir")
			memoryMode = c.Bool("memory")
		)

		api, err := getAPI(c)
		if err != nil {
			return err
		}

		// Enable memory mode.
		if memoryMode {
			d, err := drive.Open(ctx, api, resolve)
			if err != nil {
				return err
			}
			defer d.Close(ctx)

			fmt.Printf("Open drive: %s (memory mode)\n", d.Address())

			scanner := bufio.NewScanner(os.Stdout)
			for scanner.Scan() {
				args := strings.Split(scanner.Text(), " ")
				if len(args) == 0 {
					fmt.Printf(`available commands:
	- add   Add file into drive.
	- ls    List all existing files.
	- exit  Exit the cli program.
`)
					continue
				}

				cmd := args[0]
				if cmd == "add" {
					if len(args) == 2 {
						fpath := args[1]
						key := filepath.Base(fpath)
						f, err := d.Add(ctx, key, fpath)
						if err != nil {
							return err
						}
						fmt.Printf("%s (%s): %s\n", f.Key, f.Cid, humanize.IBytes(uint64(f.Size)))
					} else {
						fmt.Println("usage: add <file>")
					}
				} else if cmd == "ls" {
					prefix := ""
					if len(args) == 2 {
						prefix = args[1]
					}
					r, err := d.List(ctx, prefix)
					if err != nil {
						return err
					}

					fs := r.Files()
					for i := range fs {
						fmt.Printf("%s (%s): %s\n", fs[i].Key, fs[i].Cid, humanize.IBytes(uint64(fs[i].Size)))
					}
				} else if cmd == "exit" {
					fmt.Println("exiting...")
					return nil
				} else {
					fmt.Printf(`available commands:
	- add   Add file into drive.
	- ls    List all existing files.
	- exit  Exit the cli program.
`)
				}
			}

			return nil
		}

		if len(dir) == 0 {
			dir = defaultOrbitDBPath()
		}

		d, err := drive.Open(ctx, api, resolve, options.OpenDrive().SetDirectory(dir))
		if err != nil {
			return err
		}
		defer d.Close(ctx)

		fmt.Printf("Open drive: %s\n", d.Address())

		scanner := bufio.NewScanner(os.Stdout)
		for scanner.Scan() {
			args := strings.Split(scanner.Text(), " ")
			if len(args) == 0 {
				fmt.Println("")
			}
		}

		return nil
	},
}
