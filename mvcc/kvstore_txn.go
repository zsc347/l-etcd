package mvcc

import (
	"github.com/l-etcd/mvcc/backend"
	"github.com/l-etcd/mvcc/mvccpb"
	"github.com/l-etcd/pkg/traceutil"
)

type storeTxnRead struct {
	s  *store
	tx backend.ReadTx

	firstRev int64
	rev      int64

	trace *traceutil.Trace
}

type storeTxnWrite struct {
	storeTxnRead
	tx backend.BatchTx
	// beginRev is the revision where the txn begins;
	// it will write to the next revision.
	beginRev int64
	changes  []mvccpb.KeyValue
}

func (s *store) Write(trace *traceutil.Trace) TxnWrite {
	s.mu.RLock()
	tx := s.b.BatchTx()
	tx.Lock()

	tw := &storeTxnWrite{
		storeTxnRead: storeTxnRead{s, tx, 0, 0, trace},
		tx:           tx,
		beginRev:     s.currentRev,
		changes:      make([]mvccpb.KeyValue, 0, 4),
	}
}
