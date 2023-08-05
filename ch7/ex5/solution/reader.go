package reader

import (
	"io"
)

type reader struct {
	r        io.Reader
	n, limit int64
}

func (r *reader) Read(p []byte) (n int, err error) {
	// Fill buffer with as many bytes as limit allows using the underlying reader
	n, err = r.r.Read(p[:r.limit])
	r.n += int64(n)

	// Make sure to set the error to EOF if the limit of readable bytes has been reached
	if r.n >= r.limit {
		err = io.EOF
	}

	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &reader{r: r, limit: n}
}
