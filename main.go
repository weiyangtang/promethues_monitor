package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"prometheus_monitor/application_monitor"
	"prometheus_monitor/conf"
)

var addr = flag.String("listen-address", ":8090", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()
	conf.Init()
	application_monitor.Init()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
