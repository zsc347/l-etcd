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
)
