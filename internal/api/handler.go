package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haribabu/cardano-node-monitor/internal/monitor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Handler struct {
	monitor *monitor.Monitor
}

func NewHandler(m *monitor.Monitor) *Handler {
	return &Handler{monitor: m}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", h.healthCheck)
		v1.GET("/metrics", h.getMetrics)
		v1.GET("/status", h.getStatus)
	}

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/", h.welcome)
}

func (h *Handler) welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Cardano Node Monitor API",
		"version": "1.0.0",
		"endpoints": map[string]string{
			"health":           "/api/v1/health",
			"metrics":          "/api/v1/metrics",
			"status":           "/api/v1/status",
			"prometheus_metrics": "/metrics",
		},
	})
}

func (h *Handler) healthCheck(c *gin.Context) {
	metrics := h.monitor.GetMetrics()
	
	status := "healthy"
	if !metrics.Healthy {
		status = "unhealthy"
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":    status,
			"timestamp": metrics.Timestamp,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      status,
		"timestamp":   metrics.Timestamp,
		"uptime":      metrics.Uptime,
		"block_height": metrics.BlockHeight,
	})
}

func (h *Handler) getMetrics(c *gin.Context) {
	metrics := h.monitor.GetMetrics()
	c.JSON(http.StatusOK, metrics)
}

func (h *Handler) getStatus(c *gin.Context) {
	metrics := h.monitor.GetMetrics()
	
	syncStatus := "syncing"
	if metrics.SyncProgress >= 99.9 {
		syncStatus = "synced"
	}

	c.JSON(http.StatusOK, gin.H{
		"node_id":       metrics.NodeID,
		"healthy":       metrics.Healthy,
		"sync_status":   syncStatus,
		"sync_progress": metrics.SyncProgress,
		"block_height":  metrics.BlockHeight,
		"connections":   metrics.Connections,
		"uptime":        metrics.Uptime,
		"last_update":   metrics.Timestamp,
		"latency_ms":    metrics.Latency,
	})
}