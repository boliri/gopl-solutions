package ftp

import (
	"io/fs"
	"strings"
)

// PathError wraps an *os.PathError and strips the root of an FtpFS from the error message
type pathError struct {
	fs  FtpFS
	err *fs.PathError
}

func (pe *pathError) Error() string {
	path := strings.ReplaceAll(pe.err.Path, string(pe.fs), "")
	err := strings.ReplaceAll(pe.err.Error(), string(pe.fs), "")
	return pe.err.Op + " " + path + ": " + err
}
