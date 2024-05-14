package main

import (
	"E3DC-Prometheus-Exporter/e3dc_exporter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	exporter := e3dc_exporter.NewExporter()
	prometheus.MustRegister(exporter)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`<html>
             <head><title>E3DC Exporter</title></head>
             <body>
             <h1>Mirth Channel Exporter</h1>
             <p><a href='/metrics'>Metrics</a></p>
             </body>
             </html>`))
	})
	log.Fatal(http.ListenAndServe(":10998", nil))

}
