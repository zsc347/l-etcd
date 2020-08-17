package mvcc

type metricsTxnWrite struct {
	TxnWrite
	ranges  uint
	put     uint
	deletes uint
	putSize int64
}

func newMetricsTxnRead(tr TxnRead) TxnRead {
	return &metricsTxnWrite{&txnReadWrite{tr}, 0, 0, 0, 0}
}

func newMetricsTxnWrite(tw TxnWrite) TxnWrite {
	return &metricsTxnWrite{tw, 0, 0, 0, 0}
}
