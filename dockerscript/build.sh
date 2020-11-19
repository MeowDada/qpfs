#!/bin/sh

mkdir -p $GOPATH/src/github.com/meowdada

cd $GOPATH/src/github.com/meowdada

git clone https://github.com/meowdada/qpfs

cd qpfs

make

cp qpfs /data