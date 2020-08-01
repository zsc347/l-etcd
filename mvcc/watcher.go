package mvcc

import "github.com/l-etcd/mvcc/mvccpb"

// WatchResponse is the response for watch
type WatchResponse struct {
	// WatchID is the WatchID of the watcher this reponse sent to.
	WatchID WatchID

	// Events contains all the events that needs to send.
	Events []mvccpb.Event

	// Revision is the revision of the KV when the watchResponse is created.
	// For a normal response, the revision should be the same as the last
	// modified revision inside Events. For a delayed response to a unsynced
	// watcher, the revision is greater than the last modified revision
	// inside Events.
	Revision int64

	// CompactRevision is set when the watcher is cancelled due to compaction.
	CompactRevision int64
}

type WatchId int64

// FilterFunc returns true if the given event should be filtered out.
type filterFunc func(e mvccpb.Event) bool

type WatchStream interface {
	// Watch creates a watcher. The watcher watches the events happening or
	// happened on the given key or range [key, end] from the given startRev.
	//
	// The whole event history can be watched unless compacted.
	// If `startRev` <=0, watch observes event after currentRev.
	//
	// The return `id` is the ID of this watcher. It apperars as WatchID
	// in events that are sent to the created watcher through stream chanel.
	// The watch ID is used when it's not equal to AutoWatchID. Otherwise,
	// an auto-generated watchID is returned.
	Watch(id WatchID, key, end []byte, startRev int64, fcs ...filterFunc) (WatchID, error)

	// Chan returns a chan. All watch response will be sent to the returned chan.
	Chan() <-chan WatchResponse

	// RequestProgress requests the progress of the watcher with given ID. The response
	// will only be sent if the watcher is currently synced.
	// The responses will be sent through the WatchResponse Chan attached
	// with this stream to ensure correct ordering.
	// The responses contains no events. The revision in the response is the progress
	// of the watchers since the watcher is currently synced.
	RequestProgress(id WatchID)

	// cancel cancels a watcher by giving its ID. If watcher does not exist, an error will
	// be returned
	Cancel(id WatchID) error

	// Close closes Chan and release all related resources.
	Close()

	// Rev returns the current revision of the KV the stream watches on.
	Rev() int64
}
