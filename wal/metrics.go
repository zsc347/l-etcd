package wal

import "github.com/prometheus/client_golang/prometheus"

var (
	walFsyncSec = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "etcd",
		Subsystem: "disk",
		Name:      "wal_fsync_duration_seconds",
		Help:      "The lattency distributions of fsync called by WAL.",

		// lowest bucket start of upper bound 0.001 sec (1ms) with factor2
		// highest bucket start of 0.001 sec * 2^13 == 8.192sec
		Buckets: prometheus.ExponentialBuckets(0.001, 2, 14),
	})

	walWriteBytes = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "etcd",
		Subsystem: "disk",
		Name:      "wal_write_bytes_total",
		Help:      "Total number of bytes written in WAL.",
	})
)

func init() {
	prometheus.MustRegister(walFsyncSec)
	prometheus.MustRegister(walWriteBytes)
}
