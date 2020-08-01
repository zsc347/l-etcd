package mvcc

import "github.com/l-etcd/mvcc/mvccpb"

// WatchID is type for watch
type WatchID int64

// FilterFunc returns true if the given event should be filtered out.
type FilterFunc func(e mvccpb.Event) bool

type watcher struct {
	// the watcher key
	key []byte

	// end indicates the end of the range to watch.
	// If the end is set, the watcher is on a range.
	end []byte

	// victim is set when ch is blocked and undergoing victim processing
	victim bool

	// compacted is set when the watcher is removed because of compaction
	compacted bool

	// restore is true when the watcher is being restored from leader snapshot
	// which means that this watcher has just been moved from "synced" to "unsynced"
	// watcher group, possibly with a future revision when it was first added
	// to the synced watcher
	// "unsynced" watcher revision must always be <= current revision,
	// except when the watcher were to be moved from "synced" watcher group
	restore bool

	// minRev is the minimum revision update the watcher will accept
	minRev int64
	id     WatchID

	fcs []FilterFunc

	// a chan to send out the watch response.
	// The chan might be shared with other watchers.
	ch chan<- WatchResponse
}
