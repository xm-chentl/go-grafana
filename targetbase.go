package gografana

import "github.com/xm-chentl/go-grafana/targetformat"

type DashboardPanelTargetBase struct {
	Alias        string             // 别名
	DataSource   string             // Grafana 配置的数据源
	Query        string             // 查询语句
	RefID        string             // 多指标的别名
	Hide         bool               // 是否隐藏
	TargetFormat targetformat.Value // 结果格式

}

func (d *DashboardPanelTargetBase) SetFormat(targetFormat targetformat.Value) {
	d.TargetFormat = targetFormat
}
