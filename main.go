package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var messagesReceived = prometheus.NewDesc(
	prometheus.BuildFQName("namespace", "", "testname"),
	"HelpString",
	[]string{"channel"},
	nil,
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":10199", nil))

}
