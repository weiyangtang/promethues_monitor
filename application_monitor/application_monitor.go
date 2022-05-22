package application_monitor

import (
	"fmt"
	"github.com/spf13/viper"
	"prometheus_monitor/consts"
	"prometheus_monitor/machine_monitor"
	"prometheus_monitor/prometheus"
	"strconv"
	"time"
)

func init() {
	go ServerHealthCheck()
}

func ServerHealthCheck() {
	// check port bind ==> server is healthy
	for {
		checkPortHealthy()
		fmt.Println("finish check port health")
		time.Sleep(60 * time.Second)
	}
}

func checkPortHealthy() {
	ports := viper.GetIntSlice("health.check_ports")
	for _, port := range ports {
		ipPort := fmt.Sprintf("localhost:%v", port)
		isBind := machine_monitor.CheckPortBind(ipPort)
		portBindCount := 0
		if isBind {
			portBindCount = 1
		}
		prometheus.GaugeVecPortCount(
			float64(portBindCount),
			map[string]string{consts.MetricsTagServerPort: strconv.Itoa(port)},
		)
	}
}
