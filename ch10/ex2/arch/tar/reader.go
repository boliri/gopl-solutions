package tar

import (
	"arch"
	"fmt"
	"io"
	"strings"

	tarpkg "archive/tar"
)

const format = "tar"

// The magic number for POSIX TAR files starts at the 257th byte
var magic = strings.Repeat("?", 257) + "ustar\000"

// read reads the files in a POSIX TAR file and returns an Archive holding their filenames
func read(r io.Reader) (*arch.Archive, error) {
	tr := tarpkg.NewReader(r)

	a := &arch.Archive{}
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return a, err // return what we found so far, along with the error
		}

		a.Add(header.Name)
	}
	return a, nil
}

func init() {
	arch.RegisterFormat(format, magic, read)
	fmt.Println("TAR format registered")
}
