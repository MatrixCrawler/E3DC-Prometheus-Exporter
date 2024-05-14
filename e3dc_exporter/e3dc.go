package e3dc_exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spali/go-rscp/rscp"
	"strings"
)

type e3dcValue struct {
	Tag         rscp.Tag
	Description *prometheus.Desc
	Message     rscp.Message
}

var valuesToQuery = []e3dcValue{
	{
		Tag: rscp.EMS_POWER_PV,
		Description: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", strings.ToLower(rscp.EMS_POWER_PV.String())),
			"How much power does the pv system generate",
			[]string{"channel"}, nil,
		),
	},
	{
		Tag: rscp.EMS_POWER_BAT,
		Description: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", strings.ToLower(rscp.EMS_POWER_BAT.String())),
			"How much power does the pv system send to the battery",
			[]string{"channel"}, nil,
		),
	},
}
