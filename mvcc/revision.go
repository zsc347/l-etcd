package mvcc

import "encoding/binary"

// revBytesLen is the byte length of a normal revision.
// First 8 bytes is the revision.main in big-endian format.
// The 9th byte is a '_'.
// The last 8 bytes is the revision.sub in big-endian format
const revBytesLen = 8 + 1 + 8

// A revision indicates modification of the key-value space.
// The set of changes that share same main revision changes the
// key-value space automically.
type revision struct {
	// main is the main revision of a set of changes that happen atomically.
	main int64

	// sub is the sub revision of a change in a set of changes that happen
	// atomatically. Each change has different increasing sub revision in
	// that set.
	sub int64
}

func (a revision) GreaterThan(b revision) bool {
	if a.main > b.main {
		return true
	}
	if a.main < b.main {
		return false
	}
	return a.sub > b.sub
}

func newRevBytes() []byte {
	return make([]byte, revBytesLen, markedRevBytesLen)
}

func revToBytes(rev revision, bytes []byte) {
	binary.BigEndian.PutUint64(bytes, uint64(rev.main))
	bytes[8] = '_'
	binary.BigEndian.PutUint64(bytes[9:], uint64(rev.sub))
}

func bytesToRev(bytes []byte) revision {
	return revision{
		main: int64(binary.BigEndian.Uint64(bytes[0:8])),
		sub:  int64(binary.BigEndian.Uint64(bytes[9:])),
	}
}

type revisions []revision

func (a revisions) Len() int           { return len(a) }
func (a revisions) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a revisions) Less(i, j int) bool { return a[j].GreaterThan(a[i]) }
