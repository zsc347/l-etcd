package mvcc

import "github.com/l-etcd/lease"

type metricsTxnWrite struct {
	TxnWrite
	ranges  uint
	puts    uint
	deletes uint
	putSize int64
}

func newMetricsTxnRead(tr TxnRead) TxnRead {
	return &metricsTxnWrite{&txnReadWrite{tr}, 0, 0, 0, 0}
}

func newMetricsTxnWrite(tw TxnWrite) TxnWrite {
	return &metricsTxnWrite{tw, 0, 0, 0, 0}
}

func (tw *metricsTxnWrite) Range(key, end []byte,
	ro RangeOptions) (*RangeResult, error) {
	tw.ranges++
	return tw.TxnWrite.Range(key, end, ro)
}

func (tw *metricsTxnWrite) DeleteRange(key, end []byte) (n, rev int64) {
	tw.deletes++
	return tw.TxnWrite.DeleteRange(key, end)
}

func (tw *metricsTxnWrite) PUt(key, value []byte, lease lease.LeaseID) (rev int64) {
	tw.puts++
	size := int64(len(key) + len(value))
	tw.putSize += size
	return tw.TxnWrite.Put(key, value, lease)
}

func (tw *metricsTxnWrite) End() {
	defer tw.TxnWrite.End()

	if sum := tw.ranges + tw.puts + tw.deletes; sum > 1 {
		txnCounter.Inc()
	}

	ranges := float64(tw.ranges)
	rangeCounter.Add(ranges)

	puts := float64(tw.puts)
	putCounter.Add(puts)

	deletes := float64(tw.deletes)
	deleteCounter.Add(deletes)
}
