package util

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/dustin/go-humanize"
	"github.com/urfave/cli/v2"
)

const (
	maxBlockSize uint64 = 32 * humanize.MiByte
)

var genBlockCmd = &cli.Command{
	Name:  "genblock",
	Usage: "Generates blocks with random content",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "dir",
			Aliases: []string{"d"},
			Usage:   "Directory to place generating blocks",
		},
		&cli.BoolFlag{
			Name:    "combine",
			Aliases: []string{"c"},
			Usage:   "Combines all generating blocks into a file",
		},
		&cli.IntFlag{
			Name:    "amount",
			Aliases: []string{"n"},
			Usage:   "Number of blocks to generate",
			Value:   1,
		},
		&cli.StringFlag{
			Name:    "size",
			Aliases: []string{"s"},
			Usage:   "Size of the block",
			Value:   "256kib",
		},
	},
	Action: func(c *cli.Context) error {
		sizeStr := c.String("size")
		numBlocks := c.Int("amount")
		combine := c.Bool("combine")
		dir := c.String("dir")

		size, err := humanize.ParseBytes(sizeStr)
		if err != nil {
			return err
		}

		if size > maxBlockSize {
			return fmt.Errorf("cannot create such a large block")
		}

		if combine {
			fname := filepath.Join(dir, "blk.combine")
			f, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
			if err != nil {
				return err
			}
			defer f.Close()

			for i := 0; i < numBlocks; i++ {
				blk := newBlock(int64(size))
				_, err = io.Copy(f, bytes.NewBuffer(blk.data))
				if err != nil {
					return err
				}
			}

			return nil
		}

		for i := 0; i < numBlocks; i++ {
			blk := newBlock(int64(size))
			fname := filepath.Join(dir, fmt.Sprintf("blk.%08d", i))
			if err := blk.ToFile(fname); err != nil {
				return err
			}
		}

		return nil
	},
}
