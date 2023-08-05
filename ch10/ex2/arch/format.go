package arch

import (
	"bufio"
	"errors"
	"io"
	"sync"
	"sync/atomic"
)

var ErrUnsupportedFormat = errors.New("arch: unsupported format")

// Formats is the list of registered formats.
var (
	formatsMu     sync.Mutex
	atomicFormats atomic.Value
)

// A format holds an archive format's name, magic header and how to read its contents.
type format struct {
	name, magic string
	read        func(io.Reader) (*Archive, error)
}

// RegisterFormat registers an archive format for use by Read.
// Name is the name of the format, like "tar" or "zip".
// Magic is the magic prefix that identifies the format's encoding. The magic
// string can contain "?" wildcards that each match any one byte.
// Read is the function that reads the archive.
func RegisterFormat(name, magic string, read func(io.Reader) (*Archive, error)) {
	formatsMu.Lock()
	formats, _ := atomicFormats.Load().([]format)
	atomicFormats.Store(append(formats, format{name, magic, read}))
	formatsMu.Unlock()
}

// A reader is an io.Reader that can also peek ahead.
type reader interface {
	io.Reader
	Peek(int) ([]byte, error)
}

// asReader converts an io.Reader to a reader.
func asReader(r io.Reader) reader {
	if rr, ok := r.(reader); ok {
		return rr
	}
	return bufio.NewReader(r)
}

// Match reports whether magic matches b. Magic may contain "?" wildcards.
func match(magic string, b []byte) bool {
	if len(magic) != len(b) {
		return false
	}
	for i, c := range b {
		if magic[i] != c && magic[i] != '?' {
			return false
		}
	}
	return true
}

// Sniff determines the format of r's data.
func sniff(r reader) format {
	formats, _ := atomicFormats.Load().([]format)
	for _, f := range formats {
		b, err := r.Peek(len(f.magic))
		if err == nil && match(f.magic, b) {
			return f
		}
	}
	return format{}
}

// Read reads an archive with one of the registered formats.
// The string returned is the format name used during format registration.
// Format registration is typically done by an init function in the archive-
// specific package.
func Read(r io.Reader) (*Archive, string, error) {
	rr := asReader(r)
	f := sniff(rr)
	if f.read == nil {
		return nil, "", ErrUnsupportedFormat
	}
	a, err := f.read(rr)
	return a, f.name, err
}
