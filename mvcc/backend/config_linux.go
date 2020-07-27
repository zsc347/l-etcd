package backend

import (
	"syscall"

	bolt "go.etcd.io/bbolt"
)

// syscall.MAP_POPULATE on linux 2.6.23+ does sequential read-ahead
// which can speed up entire-database read with boltdb. We want to
// enable MAP_POPULATE for faster key-value store recovery in storage.
var boltOpenOptions = &bolt.Options{
	MmapFlags:      syscall.MAP_POPULATE,
	NoFreelistSync: true,
}

func (bcfg *BackendConfig) mmapSize() int {
	return int(bcfg.MmapSize)
}
