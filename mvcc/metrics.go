package mvcc

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	rangeCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "etcd",
			Subsystem: "mvcc",
			Name:      "range_total",
			Help:      "Total number of ranges seen by this member.",
		})

	putCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "etcd",
			Subsystem: "mvcc",
			Name:      "put_total",
			Help:      "Total number of puts seen by this member.",
		})

	deleteCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "etcd",
			Subsystem: "mvcc",
			Name:      "delete_total",
			Help:      "Total number of deletes seen by this member.",
		})

	txnCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "etcd",
			Subsystem: "mvcc",
			Name:      "txn_total",
			Help:      "Total number of txns seen by this member.",
		})
	keysGuage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "etcd_debugging",
			Subsystem: "mvcc",
			Name:      "keys_total",
			Help:      "Total number of keys.",
		})

	hashSec = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "etcd",
		Subsystem: "mvcc",
		Name:      "hash_duration_seconds",
		Help:      "The latency distribution of storage hash operation.",

		// 100MB usually takes 100 ms, so start with 10MB of 10 ms
		// lowest bucket start of upper bound 0.01 sec (10ms) with factor 2
		// highest bucket start of 0.01 sec *2^14 == 163.84 sec
		Buckets: prometheus.ExponentialBuckets(.01, 2, 15),
	})

	hashRevSec = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "ectd",
		Subsystem: "mvcc",
		Name:      "hash_rev_duration_seconds",
		Help:      "The latency distribution of storage hash by revision operation.",

		// 100MB usually takes 100ms, so start with 10MB of 10 ms
		// lowest bucket start of upper bound 0.01 sec (10ms) with factor 2
		// highest bucket start of 0.01 sec * 2^14 == 163.84 sec
		Buckets: prometheus.ExponentialBuckets(.01, 2, 15),
	})

	indexCompactionPauseMs = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "etcd_debugging",
			Subsystem: "mvcc",
			Name:      "index_compaction_pause_duration_milliseconds",
			Help:      "Bucketed histogram of index compaction pause duration.",

			// lowest bucket start of upper bound 0.5 ms with factor 2
			// highest bucket start of 0.5 ms * 2^13 == 4.096 sec
			Buckets: prometheus.ExponentialBuckets(0.5, 2, 14),
		})

	// overridden by mvcc initialization
	reportDbTotalSizeInBytesMu sync.RWMutex
	reportDbTotalSizeInBytes   = func() float64 { return 0 }

	dbTotalSize = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "etcd",
		Subsystem: "mvcc",
		Name:      "db_total_size_in_bytes",
		Help:      "Total size of the underlying database physically allocated in bytes.",
	}, func() float64 {
		reportDbTotalSizeInBytesMu.RLock()
		defer reportDbTotalSizeInBytesMu.RUnlock()
		return reportDbTotalSizeInBytes()
	})

	// overriden by mvcc initialization
	reportDbTotalSizeInUseInBytesMu sync.RWMutex
	reportDbTotalSizeInUseInBytes   = func() float64 { return 0 }
	dbTotalSizeInUse                = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "etcd",
		Subsystem: "mvcc",
		Name:      "db_total_size_in_use_in_bytes",
		Help:      "Total size of the underlying database logically in use in bytes.",
	}, func() float64 {
		reportDbTotalSizeInUseInBytesMu.RLock()
		defer reportDbTotalSizeInUseInBytesMu.RUnlock()
		return reportDbTotalSizeInUseInBytes()
	})

	// overridden by mvcc initialization
	reportDbOpenReadTxNMu sync.RWMutex
	reportDbOpenReadTxN   = func() float64 { return 0 }
	dbOpenReadTxN         = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "etcd",
		Subsystem: "mvcc",
		Name:      "db_open_read_transactions",
		Help:      "The number of currently open read transactions",
	}, func() float64 {
		reportDbOpenReadTxNMu.RLock()
		defer reportDbOpenReadTxNMu.RUnlock()
		return reportDbOpenReadTxN()
	})

	// overridden by mvcc initialization
	reportCurrentRevMu sync.RWMutex
	reportCurrentRev   = func() float64 { return 0 }
	currentRev         = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "etcd_debugging",
		Subsystem: "mvcc",
		Name:      "current_revision",
		Help:      "The current revision of store.",
	}, func() float64 {
		reportCurrentRevMu.RLock()
		defer reportCurrentRevMu.RUnlock()
		return reportCurrentRev()
	})

	// overridden by mvcc initialization
	reportCompactRevMu sync.RWMutex
	reportCompactRev   = func() float64 { return 0 }
	compactRev         = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "etcd_debugging",
		Subsystem: "mvcc",
		Name:      "compact_revision",
		Help:      "The revision of the last compaction in store.",
	}, func() float64 {
		reportCompactRevMu.RLock()
		defer reportCompactRevMu.RUnlock()
		return reportCompactRev()
	})

	totalPutSizeGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "etcd_debugging",
			Subsystem: "mvcc",
			Name:      "total_put_size_in_bytes",
			Help:      "The total size of put kv pairs seen by this member.",
		})
)

func init() {
	prometheus.MustRegister(rangeCounter)
	prometheus.MustRegister(putCounter)
	prometheus.MustRegister(deleteCounter)
	prometheus.MustRegister(txnCounter)
	prometheus.MustRegister(keysGuage)

	prometheus.MustRegister(hashSec)
	prometheus.MustRegister(hashRevSec)
	prometheus.MustRegister(indexCompactionPauseMs)

	prometheus.MustRegister(dbTotalSize)
	prometheus.MustRegister(dbOpenReadTxN)
	prometheus.MustRegister(dbOpenReadTxN)

	prometheus.MustRegister(currentRev)
	prometheus.MustRegister(compactRev)

	prometheus.MustRegister(totalPutSizeGauge)
}
