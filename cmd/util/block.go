package util

import (
	"bytes"
	"crypto/rand"
	"io"
	"os"
)

type block struct {
	data []byte
	size int64
}

func newBlock(size int64) *block {
	data := make([]byte, size)
	fillRandomData(data)
	return &block{
		data: data,
		size: size,
	}
}

func fillRandomData(buf []byte) {
	rand.Read(buf)
}

func (b *block) ToFile(path string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, bytes.NewBuffer(b.data))
	return err
}
