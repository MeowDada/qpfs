package ipfs

import (
	clnt "github.com/ipfs/go-ipfs-http-client"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	ma "github.com/multiformats/go-multiaddr"
)

// NewAPI creates api instance by given ipfs http client endpoint.
func NewAPI(addr string) (coreiface.CoreAPI, error) {
	maAddr, err := ma.NewMultiaddr(addr)
	if err != nil {
		return nil, err
	}
	return clnt.NewApi(maAddr)
}
