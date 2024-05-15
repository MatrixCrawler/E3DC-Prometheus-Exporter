package e3dc_exporter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseConfig(t *testing.T) {
	logger = InitLogger(ExporterConfig{Log: struct {
		Level  string
		Output string
		File   string
	}{Level: "DEBUG"}})
	parseConfig("../../test/testdata/config.yml")
	assert.Equal(t, "DEBUG", config.ExporterConfig.Log.Level)
	assert.Equal(t, "file", config.ExporterConfig.Log.Output)
	assert.Equal(t, "log/e3dc-exporter.log", config.ExporterConfig.Log.File)
	assert.Equal(t, "172.0.0.1", config.E3DC.Address)
	assert.Equal(t, "e3dc_user", config.E3DC.Username)
	assert.Equal(t, "supersafepassword", config.E3DC.Password)
	assert.Equal(t, "Sup3rS4v3K3y", config.E3DC.Key)
}
