package wal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/l-edcd/pkg/fileutil"
	"go.uber.org/zap"
)

type filePipeline struct {
	lg *zap.Logger

	// dir to put files
	dir string

	// size of files to make, in bytes
	size int64

	// count number of files generated
	count int

	filec chan *fileutil.LockedFile
	errc  chan error
	donec chan struct{}
}

func newFilePipeline(lg *zap.Logger, dir string, fileSize int64) *filePipeline {
	if lg == nil {
		lg = zap.NewNop()
	}

	fp := &filePipeline{
		lg:    lg,
		dir:   dir,
		size:  fileSize,
		filec: make(chan *fileutil.LockedFile),
		errc:  make(chan error, 1),
		donec: make(chan struct{}),
	}

	go fp.run()
	return fp
}

func (fp *filePipeline) alloc() (f *fileutil.LockedFile, err error) {
	// count %2 so this file isn't the same as the one last published
	fpath := filepath.Join(fp.dir, fmt.Sprintf("%d.tmp", fp.count%2))
	if f, err = fileutil.LockFile(fpath, os.O_CREATE|os.O_WRONLY,
		fileutil.PrivateFileMode); err != nil {
		return nil, err
	}

	if err = fileutil.Preallocate(f.File, fp.size, true); err != nil {
		fp.lg.Error("failed to preallocate space when creating a new WAL",
			zap.Int64("size", fp.size), zap.Error(err))
		f.Close()
		return nil, err
	}
	fp.count++
	return f, nil
}

// Open returns a fresh file for writing.
// Rename the file before calling Open again or there will be file collisions.
func (fp *filePipeline) Open() (f *fileutil.LockedFile, err error) {
	select {
	case f = <-fp.filec:
	case err = <-fp.errc:
	}
	return f, err
}

func (fp *filePipeline) Close() error {
	close(fp.donec)
	return <-fp.errc
}

func (fp *filePipeline) run() {
	defer close(fp.errc)

	for {
		f, err := fp.alloc()
		if err != nil {
			fp.errc <- err
			return
		}

		select {
		case fp.filec <- f:
		case <-fp.donec:
			os.Remove(f.Name())
			f.Close()
			return
		}
	}
}
