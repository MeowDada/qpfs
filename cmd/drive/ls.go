package drive

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

const (
	orbitdbPrefix  = "bafyrei"
	orbitdbHashLen = len("bafyreib3sej42zcltqktwza4c5mejwlmgy5la6n74q7kxkjqqafkpzng7a")
)

// ListDriveResult denotes a data structure which covers information about a local drive.
type ListDriveResult struct {
	DBName string
	DBAddr string
}

func printListDriveResult(lrs []ListDriveResult) {
	maxLen := 0
	sort.Slice(lrs, func(i, j int) bool {
		return lrs[i].DBName < lrs[j].DBName
	})

	for i := range lrs {
		if len(lrs[i].DBName) > maxLen {
			maxLen = len(lrs[i].DBName)
		}
	}

	for i := range lrs {
		format := "%-" + strconv.Itoa(maxLen) + "s %s\n"
		fmt.Printf(format, lrs[i].DBName, lrs[i].DBAddr)
	}
}

var lsCmd = &cli.Command{
	Name:  "ls",
	Usage: "List all existing local drives",
	Flags: []cli.Flag{
		dirFlag,
	},
	Action: func(c *cli.Context) error {
		root := c.String("dir")

		var lrs []ListDriveResult

		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				return nil
			}
			if strings.Contains(path, orbitdbPrefix) {
				// Find pattern: {root}/{hash}/{dbName} and trim it to
				// {hash}/{dbName}
				trimed := strings.TrimPrefix(path, filepath.Join(root)+"/")

				strs := strings.Split(trimed, "/")
				if len(strs) == 2 {
					lrs = append(lrs, ListDriveResult{
						DBAddr: filepath.Join("/orbitdb", strs[0], strs[1]),
						DBName: strs[1],
					})
				}
			}

			return nil
		})
		if err != nil {
			return err
		}

		// Print out list drive result.
		printListDriveResult(lrs)

		return nil
	},
}
