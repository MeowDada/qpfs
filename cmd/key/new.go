package key

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"os"

	"github.com/urfave/cli/v2"
)

var newCmd = &cli.Command{
	Name:  "new",
	Usage: "Generate new swarm key for building a private IPFS network",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "path",
			Aliases: []string{"o"},
			Usage:   "Generate swarm key as file to given path",
		},
	},
	Action: func(c *cli.Context) error {
		var (
			path = c.String("path")
			w    io.Writer
		)

		w = os.Stdout
		if len(path) > 0 {
			f, err := os.Create(path)
			if err != nil {
				return err
			}
			w = f
		}

		_, err := w.Write([]byte("/key/swarm/psk/1.0.0/\n"))
		if err != nil {
			return err
		}

		_, err = w.Write([]byte("/base16/\n"))
		if err != nil {
			return err
		}

		k := make([]byte, 32)
		_, err = rand.Read(k)
		if err != nil {
			return err
		}

		_, err = w.Write([]byte(hex.EncodeToString(k)))
		return err
	},
}
