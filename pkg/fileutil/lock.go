package fileutil

import (
	"errors"
	"os"
)

// Errors for os file lock
var (
	ErrLocked = errors.New("fileutil: file already locked")
)

// LockedFile used as os file lcock
type LockedFile struct {
	*os.File
}
