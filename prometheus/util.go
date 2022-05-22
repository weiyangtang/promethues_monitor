package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"prometheus_monitor/consts"
)

// 定义需要监控 Gauge 类型对象
var (
	gaugeVecMap = make(map[string]prometheus.GaugeVec)

	GaugeVecServerPortCount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: consts.MetricsServerPortCount,
		Help: "server port count",
	}, []string{consts.MetricsServerPortCount})
)

func init() {

	// Register the summary and the histogram with Prometheus's default registry.
	prometheus.MustRegister(GaugeVecServerPortCount)
	// Add Go module build info.
	prometheus.MustRegister(collectors.NewBuildInfoCollector())
}

func GaugeVecPortCount(metricsValue float64, tagIdMap map[string]string) {
	labelNames := make([]string, 0, len(tagIdMap))
	labelValues := make([]string, 0, len(tagIdMap))
	for k, v := range tagIdMap {
		labelNames = append(labelNames, k)
		labelValues = append(labelValues, v)
	}
	GaugeVecServerPortCount.WithLabelValues(labelValues...).Set(metricsValue)
}
