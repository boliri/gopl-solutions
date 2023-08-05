package ftp

import (
	"errors"
	"io/fs"
	"os"
	"strings"
)

// FtpFS represents a file system that clients can navigate through using the FTP server
//
// it's an almost exact copy of the unexported dirFS type that lives in the os package
//
// FtpFS satisfies the fs.FS and fs.StatFS interfaces
type FtpFS string

func (ftp FtpFS) Open(name string) (fs.File, error) {
	fullname, err := ftp.join(name)
	if err != nil {
		return nil, &pathError{ftp, &os.PathError{Op: "stat", Path: name, Err: err}}
	}
	f, err := os.Open(fullname)
	if err != nil {
		perr := err.(*os.PathError)
		return nil, &pathError{ftp, perr} // nil fs.File
	}
	return f, nil
}

func (ftp FtpFS) Stat(name string) (fs.FileInfo, error) {
	fullname, err := ftp.join(name)
	if err != nil {
		return nil, &pathError{ftp, &os.PathError{Op: "stat", Path: name, Err: err}}
	}
	f, err := os.Stat(fullname)
	if err != nil {
		perr := err.(*os.PathError)
		return nil, &pathError{ftp, perr}
	}
	return f, nil
}

func (ftp FtpFS) ReadFSDirectory(name string) ([]fs.DirEntry, error) {
	entries, err := fs.ReadDir(ftp, name)
	if perr, ok := err.(*os.PathError); ok {
		err = &pathError{ftp, perr}
	}
	return entries, err
}

func (ftp FtpFS) ReadFSFile(name string) ([]byte, error) {
	b, err := fs.ReadFile(ftp, name)
	if perr, ok := err.(*os.PathError); ok {
		err = &pathError{ftp, perr}
	}
	return b, err
}

// join returns the path for name in dir.
func (ftp FtpFS) join(name string) (string, error) {
	name = ftp.trimLeftSeparator(name)
	if name == "" {
		return string(ftp), nil
	}

	if ftp == "" {
		return "", errors.New("ftpserver: FtpFS with empty root")
	}
	if !fs.ValidPath(name) {
		return "", os.ErrInvalid
	}

	// safefilepath is an internal package and cannot be imported, so compatibility in some systems
	// such as Windows is not guaranteed
	/*
		name, err := safefilepath.FromFS(name)
		if err != nil {
			return "", os.ErrInvalid
		}
	*/

	if os.IsPathSeparator(ftp[len(ftp)-1]) {
		return string(ftp) + name, nil
	}
	return string(ftp) + string(os.PathSeparator) + name, nil
}

// trimLeftBackslash removes any leading path separator in path, if any
func (ftp FtpFS) trimLeftSeparator(name string) string {
	return strings.TrimLeftFunc(name, func(r rune) bool { return r == os.PathSeparator })
}
