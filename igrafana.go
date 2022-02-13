package gografana

import (
	"github.com/xm-chentl/go-grafana/charttype"
	"github.com/xm-chentl/go-grafana/datasourcetype"
	"github.com/xm-chentl/go-grafana/targetformat"
)

// IDashboard 仪表盘接口
type IDashboard interface {
	Create() (ResponseCreateOrUpdateDashboard, error)
	Delete(uid string) error
}

// IFolder 文件夹接口
type IFolder interface {
	BuildDashboard() IDashboard
}

// IGenerate 公共的生成方法
type IGenerate interface {
	Generate() (string, error)
}
type IDashboardPanelFactory interface {
	Build(datasourcetype.Value, charttype.Value) IDashboardPanel
}

// IDashboardPanel 仪表盘图表接口
type IDashboardPanel interface {
	IGenerate

	AddTarget(IDashboardPanelTarget)
	AddThreshold(...DashboardPanelThreshold)
	AddOverride(...DashboardPanelOverride)
	SetInfo(id int32, title string)
	SetYAxes(*DashboardPanelYAxes)
}

// IDashboardPanelTarget  仪表盘图表查询目标接口
type IDashboardPanelTarget interface {
	IGenerate

	SetFormat(targetformat.Value)
}

// IGrafana 对接统一的Grafana接口
type IGrafana interface {
	DeleteDashboard(uid string) error
	PreviewDashboard(RequestPreviewDashboard) string
	SaveDashboard(DashboardInfo, ...IDashboardPanel) (*ResponseCreateOrUpdateDashboard, error)
	CreateFolder(folderName string) (*ResponseCreateOrUpdateFolder, error)
	DeleteFolder(uid string) (*ResponseDeleteFolder, error)
	SaveFolder(uid string, folderName string) (*ResponseCreateOrUpdateFolder, error)
	PanelFactory() IDashboardPanelFactory
}
