package e3dc_exporter

import (
	"github.com/spali/go-rscp/rscp"
	"gopkg.in/ini.v1"
	"log"
)

const configFileName = "config.ini"

type ExporterConfig struct {
	endpoint, username, password, key string
}

func parseConfig(fileName string) rscp.ClientConfig {
	cfg, err := ini.Load(fileName)
	if err != nil {
		log.Printf("Failed to read config file: %v", err)
		return rscp.ClientConfig{}
	}
	return rscp.ClientConfig{
		Address:  cfg.Section("").Key("address").String(),
		Username: cfg.Section("").Key("username").String(),
		Password: cfg.Section("").Key("password").String(),
		Key:      cfg.Section("").Key("key").String(),
	}
}
