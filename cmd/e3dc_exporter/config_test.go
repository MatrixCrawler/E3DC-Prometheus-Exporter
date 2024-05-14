package e3dc_exporter

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_parseConfig(t *testing.T) {
	os.Args = []string{"cmd"}
	parseConfig("../../test/testdata/config.toml")
	assert.Equal(t, "ERROR", config.ExporterConfig.LogLevel)
	assert.Equal(t, "172.0.0.1", config.E3DC.Address)
	assert.Equal(t, "e3dc_user", config.E3DC.Username)
	assert.Equal(t, "supersafepassword", config.E3DC.Password)
	assert.Equal(t, "Sup3rS4v3K3y", config.E3DC.Key)
}
