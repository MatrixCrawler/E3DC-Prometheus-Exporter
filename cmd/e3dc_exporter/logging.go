package e3dc_exporter

import (
	"github.com/sirupsen/logrus"
	"github.com/spali/go-rscp/rscp"
	"os"
)

func InitLogger(config ExporterConfig) *logrus.Logger {
	newLogger := logrus.New()
	level, err := logrus.ParseLevel(config.Log.Level)
	newLogger.Out = os.Stdout
	switch config.Log.Output {
	case "file":
		file, err := os.OpenFile(config.Log.File, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			newLogger.Out = file
		} else {
			newLogger.Info("Failed to log to file, using default stdout")
		}
	}
	if err != nil {
		newLogger.Error(err)
		newLogger.Errorf("Could not parse loglevel %s. Defaulting to WARN", config.Log.Level)
		newLogger.SetLevel(logrus.WarnLevel)
	}
	newLogger.SetLevel(level)
	newLogger.ExitFunc = os.Exit
	rscp.Log = newLogger
	return newLogger
}
