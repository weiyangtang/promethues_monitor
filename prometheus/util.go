package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// 定义需要监控 Gauge 类型对象
var (
	queueSize = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "order_service_order_queue_size",
		Help: "The size of order queue",
	}, []string{"type"})

	gaugeVecMap = make(map[string]prometheus.GaugeVec)
)

func GaugeVecMetric(metricsName string, metricsValue float64, tagIdMap map[string]string) {
	labelNames := make([]string, 0, len(tagIdMap))
	labelValues := make([]string, 0, len(tagIdMap))
	for k, v := range tagIdMap {
		labelNames = append(labelNames, k)
		labelValues = append(labelValues, v)
	}
	gaugeVec := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: metricsName,
	}, labelNames)
	gaugeVec.WithLabelValues(labelValues...).Set(metricsValue)
}
