package gografana

import (
	"github.com/xm-chentl/go-grafana/charttype"
	"github.com/xm-chentl/go-grafana/datasourcetype"
)

// todo: 后续会重构一个图表服务PanelService
type panelFactory struct{}

func (p panelFactory) Build(datasourceType datasourcetype.Value, chartType charttype.Value) IDashboardPanel {
	return &dashboardPanelBase{
		chartType: chartType,
		dbType:    datasourceType,
	}
}
