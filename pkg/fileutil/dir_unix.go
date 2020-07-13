//+build !windows

package fileutil

import "os"

const (
	// PrivateDirMode grants owner to make/remove files inside directory
	PrivateDirMode = 0700
)

// OpenDir opens a directory fo syncing.
func OpenDir(path string) (*os.File, error) {
	return os.Open(path)
}
