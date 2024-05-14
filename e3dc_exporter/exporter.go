package e3dc_exporter

import (
	"github.com/spali/go-rscp/rscp"
	"log"
)

func getData() []rscp.Message {
	client, err := rscp.NewClient(parseConfig(configFileName))
	if err != nil {
		log.Fatalf("Could not create client: %v", err)
	}
	messages := []rscp.Message{
		{Tag: rscp.EMS_POWER_PV},
	}
	result, err := client.SendMultiple(messages)
	if err != nil {
		log.Fatalf("Could not send message to e3dc: %v", err)
	}
	return result
}
