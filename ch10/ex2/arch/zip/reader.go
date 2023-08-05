package zip

import (
	"arch"
	zippkg "archive/zip"
	"bytes"
	"fmt"
	"io"
)

const (
	format = "zip"
	magic  = "PK"
)

// asReaderAt converts an io.Reader to a bytes.Reader
func asBytesReader(r io.Reader) (*bytes.Reader, error) {
	if rr, ok := r.(*bytes.Reader); ok {
		return rr, nil
	}

	buf := bytes.NewBuffer([]byte{})
	_, err := io.Copy(buf, r)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(buf.Bytes()), nil
}

// read reads the files in a ZIP file and returns an Archive holding their filenames
func read(r io.Reader) (*arch.Archive, error) {
	rr, err := asBytesReader(r)
	if err != nil {
		return nil, err
	}

	zr, err := zippkg.NewReader(rr, int64(rr.Len()))
	if err != nil {
		return nil, err
	}

	a := &arch.Archive{}
	for _, zf := range zr.File {
		a.Add(zf.Name)
	}
	return a, nil
}

func init() {
	arch.RegisterFormat(format, magic, read)
	fmt.Println("ZIP format registered")
}
