package wal

import (
	"errors"
	"hash/crc32"
	"time"
)

const (
	metadataType int64 = iota + 1
	entryType
	stateType
	crcType
	snapshotType

	// warnSyncDuration is the amount of time allotted to an fsync before
	// logging a warning
	warnSyncDuration = time.Second
)

var (
	crcTable = crc32.MakeTable(crc32.Castagnoli)
)

// Error definitions for wal
var (
	ErrMaxWALEntrySizeLimitExceeded = errors.New("wal: max entry size limit exceeded")
)
