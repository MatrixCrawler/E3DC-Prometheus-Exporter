package e3dc_exporter

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/spali/go-rscp/rscp"
)

type e3dcValue struct {
	RequestTag  rscp.Tag
	ResultTag   rscp.Tag
	Description *prometheus.Desc
	Message     rscp.Message
}

var valuesToQuery = []e3dcValue{
	{
		RequestTag: rscp.EMS_REQ_POWER_PV,
		ResultTag:  rscp.EMS_POWER_PV,
		Description: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", strings.ToLower(rscp.EMS_POWER_PV.String())),
			"How much power does the pv system generate",
			nil, nil,
		),
	},
	{
		RequestTag: rscp.EMS_REQ_POWER_BAT,
		ResultTag:  rscp.EMS_POWER_BAT,
		Description: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", strings.ToLower(rscp.EMS_POWER_BAT.String())),
			"How much power does the pv system send to the battery",
			nil, nil,
		),
	},
	{
		RequestTag: rscp.EMS_REQ_POWER_HOME,
		ResultTag:  rscp.EMS_POWER_HOME,
		Description: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", strings.ToLower(rscp.EMS_POWER_HOME.String())),
			"How much power is used overall",
			nil, nil,
		),
	},
	{
		RequestTag: rscp.EMS_REQ_POWER_GRID,
		ResultTag:  rscp.EMS_POWER_GRID,
		Description: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", strings.ToLower(rscp.EMS_POWER_GRID.String())),
			"How much power is used from the grid",
			nil, nil,
		),
	},
	{
		RequestTag: rscp.EMS_REQ_BAT_SOC,
		ResultTag:  rscp.EMS_BAT_SOC,
		Description: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", strings.ToLower(rscp.EMS_BAT_SOC.String())),
			"Battery load level in percent",
			nil, nil,
		),
	},
	{
		RequestTag: rscp.EMS_REQ_POWER_ADD,
		ResultTag:  rscp.EMS_POWER_ADD,
		Description: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", strings.ToLower(rscp.EMS_POWER_ADD.String())),
			"External power source",
			nil, nil,
		),
	},
	{
		RequestTag: rscp.EMS_REQ_AUTARKY,
		ResultTag:  rscp.EMS_AUTARKY,
		Description: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", strings.ToLower(rscp.EMS_AUTARKY.String())),
			"Autarky degree in percent",
			nil, nil,
		),
	},
	{
		RequestTag: rscp.EMS_REQ_SELF_CONSUMPTION,
		ResultTag:  rscp.EMS_SELF_CONSUMPTION,
		Description: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", strings.ToLower(rscp.EMS_SELF_CONSUMPTION.String())),
			"Autarky degree in percent",
			nil, nil,
		),
	},
}
