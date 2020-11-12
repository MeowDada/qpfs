package drive

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

var rmCmd = &cli.Command{
	Name:      "rm",
	Usage:     "Remove local drive data",
	UsageText: "qpfs drive rm [drive...]",
	Flags: []cli.Flag{
		dirFlag,
	},
	Before: func(c *cli.Context) error {
		return nil
	},
	Action: func(c *cli.Context) error {
		dir := c.String("dir")
		args := c.Args().Slice()

		// Scan the whole directory.
		drives, err := scanDrives(dir)
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(os.Stdout)
		for _, driveName := range args {
			sets, ok := drives[driveName]
			if !ok {
				fmt.Printf("There is no such drive: %s\n", driveName)
				continue
			}

			fmt.Printf("Are you sure to remove drive %s? (y/n)\n", driveName)
			for scanner.Scan() {
				text := scanner.Text()
				if text == "y" {
					for _, set := range sets {
						fmt.Printf("Is %s the one you want to remove?\n", set.GetPath(dir))
						for scanner.Scan() {
							text = scanner.Text()
							if text == "y" {
								fmt.Println("remove", set.GetPath(dir))
								if err := os.RemoveAll(set.GetPath(dir)); err != nil {
									fmt.Printf("remove with error: %v\n", err)
								}
								break
							} else if text == "n" {
								fmt.Println("Not to remove", set.GetPath(dir))
								break
							} else {
								fmt.Println(`Press 'y'to confirm or 'n' to cancel`)
							}
						}
					}
					break
				} else if text == "n" {
					fmt.Println("Not to remove", driveName)
					break
				} else {
					fmt.Println(`Press 'y'to confirm or 'n' to cancel`)
				}
			}
		}

		return nil
	},
}

func scanDrives(root string) (map[string][]driveLocation, error) {
	r := map[string][]driveLocation{}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}
		if strings.Contains(path, orbitdbPrefix) {
			trimed := strings.TrimPrefix(path, filepath.Join(root)+"/")
			strs := strings.Split(trimed, "/")

			if len(strs) == 2 {
				dbHash := strs[0]
				dbName := strs[1]
				r[dbName] = append(r[dbName], driveLocation{
					Name: dbName,
					Hash: dbHash,
				})
				return nil
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return r, nil
}
