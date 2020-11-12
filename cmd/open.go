package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/meowdada/ipfstor/drive"
	"github.com/meowdada/ipfstor/options"
	"github.com/meowdada/qpfs/cmd/shell"
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
		&cli.BoolFlag{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "Create new database if it does not present",
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
			ctx           = c.Context
			resolve       = c.Args().First()
			dir           = c.String("dir")
			createNew     = c.Bool("new")
			enableMemMode = c.Bool("memory")
		)

		if enableMemMode {
			dir = ""
		}

		api, err := getAPI(c)
		if err != nil {
			return err
		}

		opts := options.OpenDrive().
			SetDirectory(dir).
			SetCreate(createNew)

		d, err := drive.Open(ctx, api, resolve, opts)
		if err != nil {
			return err
		}
		defer d.Close(ctx)

		app := shell.App
		app.Set("drive", d)

		scanner := bufio.NewScanner(os.Stdout)

		fmt.Printf("> ")
		for scanner.Scan() {
			text := scanner.Text()
			args := strings.Split(text, " ")
			err := app.Run(ctx, args)
			if err == shell.ErrExitShell {
				fmt.Println("exiting...")
				return nil
			}
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("> ")
		}

		return nil
	},
}

var shellHelpText = `Available commands:
	- add    Add local file to current drive.
	- ls     List all existing files of current drive.
	- rm     Remove file from current drive.
	- exit   Exit this interactive shell.
`
