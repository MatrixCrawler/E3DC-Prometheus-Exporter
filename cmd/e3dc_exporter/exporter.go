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
	ch <- prometheus.MustNewConstMetric(up, prometheus.GaugeValue, 1)
	for _, value := range exporterData {
		v := getValueFromMessage(value.Message)
		ch <- prometheus.MustNewConstMetric(value.Description, prometheus.GaugeValue, v)
	}
}

func getValueFromMessage(message rscp.Message) float64 {
	switch message.DataType {
	case rscp.Char8:
		return float64(message.Value.(int8))
	case rscp.UChar8:
		return float64(message.Value.(uint8))
	case rscp.Int32:
		return float64(message.Value.(int32))
	case rscp.Uint32:
		return float64(message.Value.(uint32))
	}
	logger.Warn("Got no value from message")
	return 0
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
		messages = append(messages, rscp.Message{Tag: value.RequestTag, DataType: value.RequestTag.DataType()})
	}
	results, err := client.SendMultiple(messages)
	if err != nil {
		logger.Errorf("Could not send messages to e3dc: %v", err)
		return nil, err
	}
	var e3dcResult []e3dcValue
	for _, result := range results {
		for _, value := range e3dcValues {
			if value.ResultTag == result.Tag {
				e3dcResult = append(e3dcResult, e3dcValue{
					RequestTag:  value.RequestTag,
					ResultTag:   value.ResultTag,
					Description: value.Description,
					Message: rscp.Message{
						Tag:      result.Tag,
						DataType: result.DataType,
						Value:    result.Value,
					},
				})
			}
		}
	}
	return e3dcResult, nil
}

func getClient() (*rscp.Client, error) {
	client, err := rscp.NewClient(config.ClientConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}
