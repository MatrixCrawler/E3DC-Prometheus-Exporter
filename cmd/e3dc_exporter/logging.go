package e3dc_exporter

import (
	"github.com/sirupsen/logrus"
	"github.com/spali/go-rscp/rscp"
	"os"
)

func InitLogger(logLevel string) *logrus.Logger {
	newLogger := logrus.New()
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		newLogger.Error(err)
		newLogger.Errorf("Could not parse loglevel %s. Defaulting to WARN", logLevel)
		newLogger.SetLevel(logrus.WarnLevel)
	}
	newLogger.SetLevel(level)
	newLogger.ExitFunc = os.Exit
	rscp.Log = newLogger
	return newLogger
}
