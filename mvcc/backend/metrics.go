package backend

import "github.com/prometheus/client_golang/prometheus"

var (
	commitSec = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "etcd",
		Subsystem: "disk",
		Name:      "backend_commit_duration_seconds",
		Help:      "The latency distributions of commit called by backend.",

		// lowest bucket start of upper bound 0.001 sec (1ms) with factor 2
		// highest bucket start of 0.001 sec * 2^13 == 8.192 sec
		Buckets: prometheus.ExponentialBuckets(0.001, 2, 14),
	})

	rebalanceSec = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "etcd_debugging",
		Subsystem: "disk",
		Name:      "backend_commit_relalance_duration_seconds",
		Help:      "The latency distributions of commit.rebalance called by bboltdb backend",

		// lowest bucket start of upper bound 0.001 sec (1 ms) with factor 2
		// highest bucket start of 0.001 sec * 2^13 == 8.192 sec
		Buckets: prometheus.ExponentialBuckets(0.001, 2, 14),
	})

	spillSec = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "etcd_debugging",
		Subsystem: "disk",
		Name:      "backend_commit_spill_duration_seconds",
		Help:      "The latency distributions of commit.spill called by bboltdb backend",

		// lowest bucket start of upper bound 0.001 sec (1 ms) with factor 2
		// highest bucket start of 0.001 sec * 2^13 == 8.192 sec
		Buckets: prometheus.ExponentialBuckets(0.001, 2, 14),
	})

	writeSec = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "etcd_debugging",
		Subsystem: "disk",
		Name:      "backend_commit_write_duration_seconds",
		Help:      "The latency distributions of commit.write called by bboltdb backend.",

		// lowest bucket start of upper bound 0.001 sec (1 ms) with factor 2
		// highest bucket start of 0.001 sec * 2^13 == 8.192 sec
		Buckets: prometheus.ExponentialBuckets(0.001, 2, 14),
	})

	defragSec = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "etcd",
		Subsystem: "disk",
		Name:      "backend_defrag_duration_seconds",
		Help:      "The latency distribution of backend defragmentation.",

		// 100 MB usually takes sec, so start with 10MB of 100ms
		// lowest bucket start of upper bound 0.1 sec (100ms) with factor 2
		// highest bucket start of 0.1 sec * 2^12 == 409.6 sec
		Buckets: prometheus.ExponentialBuckets(.1, 2, 13),
	})

	snapshotTransferSec = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "etcd",
		Subsystem: "disk",
		Name:      "backend_snapshot_duration_seconds",
		Help:      "The latency distribution of backend snapshots.",

		// lowest bucket start of upper bound 0.01 sec (10ms) with factor 2
		// highest bucket start of 0.01 sec * 2^16 == 655.36 sec
		Buckets: prometheus.ExponentialBuckets(.01, 2, 17),
	})
)

func init() {
	prometheus.MustRegister(commitSec)
	prometheus.MustRegister(rebalanceSec)
	prometheus.MustRegister(spillSec)
	prometheus.MustRegister(writeSec)
	prometheus.MustRegister(defragSec)
	prometheus.MustRegister(snapshotTransferSec)
}
