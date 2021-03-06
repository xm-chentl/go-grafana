## Grafana SDK ##

根据Grafana Api相关文档封装
### 接口说明

```go
// 文件所在位置：grafanaex/igrafana.go

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
	Build(dbtype.Value, charttype.Value) IDashboardPanel
}

// IDashboardPanel 仪表盘图表接口
type IDashboardPanel interface {
	IGenerate

	AddTarget(IDashboardPanelTarget)
	AddThreshold(...DashboardPanelThreshold)
	AddOverride(...DashboardPanelOverride)
	SetInfo(id int32, title string)
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
```

### 示例

```go

grafanaInst := gografana.New(opt, )
```
