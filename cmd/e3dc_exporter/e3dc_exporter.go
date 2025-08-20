package e3dc_exporter

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

const fileName = "config.yml"

var logger *logrus.Logger

func Run() {
	logger = InitLogger(ExporterConfig{Log: struct {
		Level  string
		Output string
		File   string
	}{Level: "ERROR"}})
	parseConfig(fileName)
	logger = InitLogger(config.ExporterConfig)
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
             <h1>E3DC Exporter</h1>
             <p><a href='/metrics'>Metrics</a></p>
             </body>
             </html>`))
	})
	logger.Infof("Starting e3dc exporter on port %v", config.ExporterConfig.Port)
	logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.ExporterConfig.Port), nil))
}
