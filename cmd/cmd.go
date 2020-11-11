package cmd

import (
	"io"
	"log"
	"os"
	"sort"

	coreiface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/meowdada/ipfstor/ipfsutil"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var mainApp = &cli.App{
	Name:  "qpfs",
	Usage: "A cli tool to interact with ipfs drive",
	Commands: []*cli.Command{
		openCmd,
		addCmd,
		grantCmd,
		revokeCmd,
		lsCmd,
	},
}

var apiFlag = &cli.StringFlag{
	Name:  "api",
	Usage: "Endpoint `ADDRESS` to ipfs http client",
	Value: ipfsutil.DefaultAPIAddress,
}

var dirFlag = &cli.StringFlag{
	Name:    "dir",
	Aliases: []string{"d"},
	Usage:   "`DIR` to the local datastore",
	Value:   defaultOrbitDBPath(),
}

// Main is the entrypoint
func Main(ver string) {
	mainApp.Version = ver
	if err := mainApp.Run(os.Args); err != nil {
		log.Println(err)
		return
	}
}

// Register registers multiple commands to main app.
func Register(cmds ...*cli.Command) {
	sort.Slice(cmds, func(i, j int) bool {
		return cmds[i].Name < cmds[j].Name
	})
	mainApp.Commands = append(mainApp.Commands, cmds...)
}

func setupZapLogger(w io.Writer, level zapcore.Level) *zap.Logger {
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	writeSyncer := zapcore.AddSync(w)
	core := zapcore.NewCore(encoder, writeSyncer, level)

	logger := zap.New(core)
	return logger
}

func getAPI(c *cli.Context) (coreiface.CoreAPI, error) {
	addr := c.String("api")
	return ipfsutil.NewAPI(addr)
}
