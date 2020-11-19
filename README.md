# Introduction
qpfs is a command line interface to control private team drives backed by IPFS.

>To run the cli, you will need a running IPFS node.

# Usage
## Prerequirement
Make sure that you have a running IPFS node already.

If not, try using below official IPFS image:

```bash
docker run -d \
    --name <containerName> \
    -v <ipfsStaging>:/export \
    -v <ipfsRepo>:/data/ipfs \
    -p 4001:4001 \
    -p 127.0.0.1:8080:8080 \
    -p 127.0.0.1:5001:5001 \
    ipfs/go-ipfs:latest
```

## Build
Install the qpfs binary.
```bash
go install -ldflags "-X main.version=`git rev-parse HEAD`"
```