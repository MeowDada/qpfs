package util

import (
	"bytes"
	"crypto/rand"
	"io"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/urfave/cli/v2"
)

var genFileCmd = &cli.Command{
	Name:  "genfile",
	Usage: "generate a file with some duplication blocks",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "path",
			Aliases: []string{"o"},
			Usage:   "Path to generating file",
			Value:   "file.gen",
		},
		&cli.StringFlag{
			Name:    "size",
			Aliases: []string{"s"},
			Usage:   "Size of the generating file",
			Value:   "32mib",
		},
		&cli.StringFlag{
			Name:    "blocksize",
			Aliases: []string{"b"},
			Usage:   "Size of the each underlying block",
			Value:   "256kib",
		},
		&cli.IntFlag{
			Name:    "ratio",
			Aliases: []string{"r"},
			Usage:   "Estimated deduplication ratio",
			Value:   50,
		},
	},
	Action: func(c *cli.Context) error {
		var (
			ratio        = c.Int("ratio")
			sizeStr      = c.String("size")
			blockSizeStr = c.String("blocksize")
			path         = c.String("path")
		)

		size, err := humanize.ParseBytes(sizeStr)
		if err != nil {
			return err
		}

		blockSize, err := humanize.ParseBytes(blockSizeStr)
		if err != nil {
			return err
		}

		// Generate duplicate parts.
		dupSize := size * uint64(ratio) / uint64(100)
		blk := newBlock(int64(blockSize))
		numBlks := dupSize / blockSize
		dupSize = blockSize * numBlks

		f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		buf := make([]byte, blockSize)
		for i := uint64(0); i < numBlks; i++ {
			_, err = io.CopyBuffer(f, bytes.NewBuffer(blk.data), buf)
			if err != nil {
				return err
			}
		}

		// Generate non-duplicate parts.
		nonDupSize := size - dupSize
		batchSize := humanize.MiByte
		numBatches := nonDupSize / uint64(batchSize)

		for i := uint64(0); i < numBatches; i++ {
			blk := newBlock(int64(batchSize))
			_, err = io.CopyBuffer(f, bytes.NewBuffer(blk.data), buf)
			if err != nil {
				return err
			}
		}

		finalPartSize := nonDupSize - (uint64(batchSize) * numBatches)
		finalPart := make([]byte, finalPartSize)

		_, err = rand.Read(finalPart)
		if err != nil {
			return err
		}

		_, err = io.CopyBuffer(f, bytes.NewBuffer(finalPart), buf)
		return err
	},
}
