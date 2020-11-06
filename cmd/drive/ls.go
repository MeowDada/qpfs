package drive

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
)

const (
	orbitdbDBPrefix = "/bafyrei"
)

var (
	home string
)

// ListDriveInfo denotes a data structure that represents a drive.
type ListDriveInfo struct {
	Addr string
	Name string
}

func defaultOrbitdbPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".orbitdb")
}

var lsCmd = &cli.Command{
	Name:        "ls",
	Usage:       "List all existing drives.",
	UsageText:   "qpfs drive ls [prefix (optional)]",
	Description: "List all existing drives with given prefix.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "dir",
			Aliases: []string{"d"},
			Usage:   "Directoy to ipfs share drives",
			Value:   defaultOrbitdbPath(),
		},
	},
	Before: func(c *cli.Context) error {
		if c.Args().Len() > 1 {
			return fmt.Errorf("cannot specific more than one prefix")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		var prefix string
		if c.Args().Len() == 1 {
			prefix = c.Args().First()
		}
		dir := c.String("dir")
		var infos []ListDriveInfo

		if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			trimed := strings.TrimPrefix(path, dir)
			if !strings.HasPrefix(trimed, orbitdbDBPrefix) {
				return nil
			}
			strs := strings.Split(trimed, "/")
			if len(strs) != 3 || !info.IsDir() {
				return nil
			}
			infos = append(infos, ListDriveInfo{
				Addr: strs[1],
				Name: strs[2],
			})
			return nil
		}); err != nil {
			return err
		}

		for i := range infos {
			if strings.HasPrefix(infos[i].Name, prefix) {
				fmt.Printf("[%s]: %s\n", infos[i].Addr, infos[i].Name)
			}
		}

		return nil
	},
}
