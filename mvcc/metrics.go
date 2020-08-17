package mvcc

import "github.com/prometheus/client_golang/prometheus"

var (
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
)
