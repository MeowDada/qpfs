package ipfs

import (
	"fmt"
	"path/filepath"
	"strings"
)

// DstPath denotes
type DstPath interface {
	Drive() string
	Filename() string
}

type absoultePath struct {
	drive string
	fname string
}

func (p *absoultePath) Drive() string {
	return p.drive
}

func (p *absoultePath) Filename() string {
	return p.fname
}

// NewDstPath try parsing input into an instance of DstPath.
func NewDstPath(src, dst string) (DstPath, error) {
	// Try Parsing dst as pattern <kv>/<fname>.
	splits := strings.Split(dst, "/")
	if len(splits) == 1 {
		return &absoultePath{
			drive: splits[0],
			fname: filepath.Base(src),
		}, nil
	}
	if len(splits) == 2 {
		return &absoultePath{
			drive: splits[0],
			fname: splits[1],
		}, nil
	}
	return nil, fmt.Errorf("invalid dst: %s", dst)
}

// ParseDriveObject parses drive and object from the given string.
func ParseDriveObject(str string) (drive string, obj string, err error) {
	splits := strings.Split(str, "/")
	if len(splits) != 2 {
		return drive, obj, fmt.Errorf("invalid target path")
	}
	return splits[0], splits[1], nil
}
