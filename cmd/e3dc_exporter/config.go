package e3dc_exporter

import (
	"github.com/pelletier/go-toml/v2"
	"github.com/spali/go-rscp/rscp"
	"os"
)

type Config struct {
	ExporterConfig ExporterConfig
	E3DC           rscp.ClientConfig
}

type ExporterConfig struct {
	LogLevel string
}

var config = Config{
	ExporterConfig: ExporterConfig{
		LogLevel: "ERROR",
	},
	E3DC: rscp.ClientConfig{},
}

func parseConfig(fileName string) rscp.ClientConfig {
	configFile, err := os.ReadFile(fileName)
	if err != nil {
		logger.Fatalf("Error opening config file %q. %v", fileName, err)
	}
	err = toml.Unmarshal(configFile, &config)
	if err != nil {
		logger.Fatalf("Error unmarshalling toml config file: %v", err)
	}
	return rscp.ClientConfig{}
}
