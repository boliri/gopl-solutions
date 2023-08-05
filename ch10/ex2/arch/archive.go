// Package arch implements a library for reading POSIX tar / zip files without decompression.
//
// It is heavily inspired by Go's image package.
package arch

import (
	"strings"
)

// Archive holds the list of filenames in a compressed file
type Archive struct {
	filenames []string
}

// Add adds a filename to the archive
func (a *Archive) Add(filename string) {
	a.filenames = append(a.filenames, filename)
}

func (a *Archive) String() string {
	if len(a.filenames) == 0 {
		return "<empty>"
	}

	return strings.Join(a.filenames, "\n")
}
