package reader

import (
	"io"
)

type reader struct {
	s string
}

func (r *reader) Read(p []byte) (n int, err error) {
	// Fill the bytes buffer with as many runes as its capacity allows
	n = copy(p, r.s)

	// Strip read bytes from the reader's string
	r.s = r.s[n:]

	if len(r.s) == 0 {
		err = io.EOF
	}

	return
}

func NewReader(s string) io.Reader {
	return &reader{s}
}
