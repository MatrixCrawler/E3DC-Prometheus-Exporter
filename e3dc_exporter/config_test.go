package e3dc_exporter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseConfig(t *testing.T) {
	config := parseConfig("../test-data/config.ini")
	assert.Equal(t, "172.0.0.1", config.Address)
	assert.Equal(t, "e3dc_user", config.Username)
	assert.Equal(t, "supersafepassword", config.Password)
	assert.Equal(t, "Sup3rS4v3K3y", config.Key)
}
