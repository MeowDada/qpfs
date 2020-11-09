package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/meowdada/ipfstor"
	"github.com/meowdada/qpfs/pkg/ipfs"
	cli "github.com/urfave/cli/v2"
)

var openCmd = &cli.Command{
	Name:  "open",
	Usage: "Open an existing drive or create a new one",
	Before: func(c *cli.Context) error {
		if c.Args().Len() != 1 {
			return fmt.Errorf("usage: qpfs open <addr>")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		ctx := c.Context
		addr := "/ip4/127.0.0.1/tcp/5001"

		api, err := ipfs.NewAPI(addr)
		if err != nil {
			return err
		}

		drive, err := ipfstor.OpenDrive(ctx, api, addr)
		if err != nil {
			return err
		}

		cmds := map[string]struct{}{
			"add": {},
			"get": {},
			"ls":  {},
			"rm":  {},
		}

		scanner := bufio.NewScanner(os.Stdout)
		fmt.Println("Enter commands: (add, get, ls, rm)")

		for scanner.Scan() {
			text := scanner.Text()
			args := strings.Split(text, " ")
			fmt.Println(args)
			if len(args) < 1 {
				return fmt.Errorf("available commands: (add, get, ls, rm)")
			}

			cmd := args[0]

			if _, ok := cmds[cmd]; !ok {
				fmt.Printf("unrecognizable command: %s\n", cmd)
				continue
			}

			if cmd == "add" {
				if len(args) != 3 {
					fmt.Println("invalid input")
					continue
				}
				if err := drive.Add(ctx, args[2], args[1]); err != nil {
					return err
				}
			} else if cmd == "ls" {
				r, err := drive.List(ctx)
				if err != nil {
					return err
				}
				fmt.Println(r)
			}
		}

		return nil
	},
}
