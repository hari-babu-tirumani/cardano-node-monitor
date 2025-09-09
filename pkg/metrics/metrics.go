package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	NodeHeight = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cardano_node_block_height",
			Help: "Current block height of the Cardano node",
		},
		[]string{"node_id"},
	)

	NodeSyncProgress = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cardano_node_sync_progress",
			Help: "Sync progress percentage of the Cardano node",
		},
		[]string{"node_id"},
	)

	NodeConnections = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cardano_node_connections",
			Help: "Number of peer connections",
		},
		[]string{"node_id"},
	)

	NodeUptime = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cardano_node_uptime_seconds",
			Help: "Node uptime in seconds",
		},
		[]string{"node_id"},
	)

	NodeLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "cardano_node_response_time_seconds",
			Help:    "Response time for node API calls",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"node_id", "endpoint"},
	)

	NodeErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cardano_node_errors_total",
			Help: "Total number of node errors",
		},
		[]string{"node_id", "error_type"},
	)
)

type NodeMetrics struct {
	NodeID       string    `json:"node_id"`
	BlockHeight  uint64    `json:"block_height"`
	SyncProgress float64   `json:"sync_progress"`
	Connections  int       `json:"connections"`
	Uptime       int64     `json:"uptime"`
	Timestamp    time.Time `json:"timestamp"`
	Healthy      bool      `json:"healthy"`
	Latency      float64   `json:"latency_ms"`
}

func RecordMetrics(metrics *NodeMetrics) {
	NodeHeight.WithLabelValues(metrics.NodeID).Set(float64(metrics.BlockHeight))
	NodeSyncProgress.WithLabelValues(metrics.NodeID).Set(metrics.SyncProgress)
	NodeConnections.WithLabelValues(metrics.NodeID).Set(float64(metrics.Connections))
	NodeUptime.WithLabelValues(metrics.NodeID).Set(float64(metrics.Uptime))
	NodeLatency.WithLabelValues(metrics.NodeID, "tip").Observe(metrics.Latency / 1000)
}