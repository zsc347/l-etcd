package fileutil

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	// PrivateFileMode grants owner to read/write a file.
	PrivateFileMode = 0600
)

// IsDirWriteable checks if dir is writable by writing and removing a file
// to dir. It returns nil if dir is writable.
func IsDirWriteable(dir string) error {
	f := filepath.Join(dir, ".touch")
	if err := ioutil.WriteFile(f, []byte(""), PrivateFileMode); err != nil {
		return err
	}
	return os.Remove(f)
}

// TouchDirAll is similar to os.MkDirAll. It creates directories with 0700
// permission if any directory does not exists. TouchDirAll also ensures the
// given directory is writable.
func TouchDirAll(dir string) error {
	// If path is already a directory, MkdirAll does nothing and returns nil.
	// So, first check if dir exist with an expected permission mode.
	if Exist(dir) {
		err := CheckDirPermission(dir, PrivateFileMode)
		if err != nil {
			return err
		}
		return IsDirWriteable(dir)
	}

	err := os.MkdirAll(dir, PrivateDirMode)
	if err != nil {
		// if mkdirAll("a/text") and "text" is not a directoy,
		// this will return syscanl.ENOTDIR
		return err
	}
	return IsDirWriteable(dir)
}

// CreateDirAll is similar to TouchDirAll but returns error
// if the deepest diretory was not empty
func CreateDirAll(dir string) error {
	err := TouchDirAll(dir)
	if err != nil {
		return err
	}

	var ns []string
	ns, err = ReadDir(dir)
	if err != nil {
		return err
	}
	if len(ns) != 0 {
		err = fmt.Errorf("expected %q to be empty, got %q", dir, ns)
	}
	return err
}

// Exist returns true if a file or directory exists.
func Exist(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

// DirEmpty returns true is a directory empty and can access.
func DirEmpty(name string) bool {
	ns, err := ReadDir(name)
	return err == nil && len(ns) == 0
}

// ZeroToEnd zeros a file starting from SEEK_CUR to its SEEK_END.
// May temporarily shoten the length of the file.
func ZeroToEnd(f *os.File) error {
	// TODO: support FALLOC_FL_ZERO_RANGE
	off, err := f.Seek(0, io.SeekCurrent)
	if err != nil {
		return err
	}

	lenf, lerr := f.Seek(0, io.SeekEnd)
	if lerr != nil {
		return lerr
	}

	if err = f.Truncate(off); err != nil {
		return err
	}

	// make sure blocks remain allocated
	if err = Preallocate(f, lenf, true); err != nil {
		return err
	}

	_, err = f.Seek(off, io.SeekStart)
	return err
}

// CheckDirPermission checks permission on an existing dir.
// Returns error if dir is empty or exist with a different permission than specified
func CheckDirPermission(dir string, perm os.FileMode) error {
	if !Exist(dir) {
		return fmt.Errorf("directory %q empty, can not check permission", dir)
	}

	// check the existing permission on the directory
	dirInfo, err := os.Stat(dir)
	if err != nil {
		return err
	}

	dirMode := dirInfo.Mode().Perm()
	if dirMode != perm {
		err = fmt.Errorf("directory %q,%q exist without desired file permission %q",
			dir, dirInfo.Mode(), os.FileMode(PrivateFileMode))
		return err
	}
	return nil
}
