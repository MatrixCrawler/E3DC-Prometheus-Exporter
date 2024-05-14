package e3dc_exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spali/go-rscp/rscp"
)

const namespace = "e3dc"

var up = prometheus.NewDesc(
	prometheus.BuildFQName(namespace, "", "up"),
	"Was the last E3DC query successful.",
	nil, nil,
)

type Exporter struct {
}

func NewExporter() *Exporter {
	return &Exporter{}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- up
	for _, value := range valuesToQuery {
		ch <- value.Description
	}
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	exporterData, err := getValues(valuesToQuery)
	if err != nil {
		ch <- prometheus.MustNewConstMetric(up, prometheus.GaugeValue, 0)
		logger.Error(err)
		return
	}
	for _, value := range exporterData {
		ch <- prometheus.MustNewConstMetric(value.Description, prometheus.GaugeValue, float64(value.Message.Value.(int32)))
	}
}

func getValues(e3dcValues []e3dcValue) ([]e3dcValue, error) {
	client, err := getClient()
	if err != nil {
		return []e3dcValue{}, err
	}
	defer func(client *rscp.Client) {
		_ = client.Disconnect()
	}(client)
	var messages []rscp.Message
	for _, value := range e3dcValues {
		messages = append(messages, rscp.Message{Tag: value.Tag})
	}
	results, err := client.SendMultiple(messages)
	if err != nil {
		logger.Errorf("Could not send messages to e3dc: %v", err)
		return nil, err
	}
	for _, result := range results {
		for _, value := range e3dcValues {
			if value.Tag == result.Tag {
				value.Message = result
			}
		}
	}
	return e3dcValues, nil
}

func getClient() (*rscp.Client, error) {
	client, err := rscp.NewClient(config.E3DC)
	if err != nil {
		return nil, err
	}
	return client, nil
}
