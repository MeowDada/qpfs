package drive

import (
	"context"

	"github.com/ipfs/go-cid"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/meowdada/ipfstor"
	"github.com/meowdada/qpfs/pkg/ipfs"
)

// API provides function utilities to interacts with a cloud drive.
type API interface {
	Add(ctx context.Context, key, fpath string) error
	Get(ctx context.Context, key string) error
	List(ctx context.Context, prefix string) (ListResult, error)
	Stop(ctx context.Context) error
}

type ListResult struct {
	files []File
}

type File struct {
	Key  string
	Addr string
	Size int64
}

// Instance denotes a drive instance backed by orbitdb.
type Instance struct {
	api   coreiface.CoreAPI
	drive ipfstor.Driver
}

// New creates an instance.
func New(ctx context.Context, dbAddr string) (*Instance, error) {
	api, err := ipfs.NewAPI("/ip4/127.0.0.1/tcp/5001")
	if err != nil {
		return nil, err
	}

	drive, err := ipfstor.NewDriver(ctx, api, dbAddr)
	return &Instance{
		api:   api,
		drive: drive,
	}, err
}

// Add implements API interface.
func (ins *Instance) Add(ctx context.Context, key, fpath string) error {
	return ins.drive.Add(ctx, key, fpath)
}

func (ins *Instance) List(ctx context.Context, prefix string) (ListResult, error) {
	var files []File

	err := ins.drive.Iter(ctx, func(ctx context.Context, key string, cid cid.Cid, size int64, owner string) error {
		files = append(files, File{
			Key:  key,
			Addr: cid.String(),
			Size: size,
		})
		return nil
	})
	return ListResult{
		files: files,
	}, err
}

func boolPtr(flag bool) *bool {
	return &flag
}

func strPtr(str string) *string {
	return &str
}
