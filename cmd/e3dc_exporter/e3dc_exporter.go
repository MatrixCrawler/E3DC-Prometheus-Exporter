package e3dc_exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
)

const fileName = "config.toml"

var logger *logrus.Logger
var version string

func Run() {
	logger = InitLogger("ERROR")
	parseConfig(fileName)
	logger = InitLogger(config.ExporterConfig.LogLevel)
	startPrometheusExporter()
}

func startPrometheusExporter() {
	exporter := NewExporter()
	logger.Debug("Registering Exporter")
	prometheus.MustRegister(exporter)
	logger.Debug("Registering /metrics handler")
	http.Handle("/metrics", promhttp.Handler())
	logger.Debug("Registering / handler")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`<html>
             <head><title>E3DC Exporter</title></head>
             <body>
             <h1>Mirth Channel Exporter</h1>
             <p><a href='/metrics'>Metrics</a></p>
             </body>
             </html>`))
	})
	logger.Info("Starting prometheus exporter")
	logger.Fatal(http.ListenAndServe(":10998", nil))
}
