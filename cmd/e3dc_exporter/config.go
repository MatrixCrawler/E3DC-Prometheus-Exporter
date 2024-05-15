package e3dc_exporter

import (
	"github.com/spali/go-rscp/rscp"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	ExporterConfig ExporterConfig
	E3DC           E3DCConfig
	ClientConfig   rscp.ClientConfig
}

type ExporterConfig struct {
	Port int
	Log  struct {
		Level  string
		Output string
		File   string
	}
}

type E3DCConfig struct {
	Address  string
	Username string
	Password string
	Key      string
}

var config = Config{}

func parseConfig(fileName string) rscp.ClientConfig {
	configFile, err := os.ReadFile(fileName)
	if err != nil {
		logger.Fatalf("Error opening config file %q. %v", fileName, err)
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		logger.Fatalf("Error unmarshalling yaml config file: %v", err)
	}
	config.ClientConfig = rscp.ClientConfig{
		Address:  config.E3DC.Address,
		Key:      config.E3DC.Key,
		Username: config.E3DC.Username,
		Password: config.E3DC.Password,
	}
	return rscp.ClientConfig{}
}
