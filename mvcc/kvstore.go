package mvcc

const (
	// markedRevBytesLen is the byte length of marked revision.
	// The first `revBytesLen` bytes represents a normal revision.
	// The last one byte is the mark.
	markedRevBytesLen = revBytesLen + 1
)
