package gografana

import (
	"bytes"
	"errors"
	"text/template"

	"github.com/xm-chentl/go-grafana/charttype"
	"github.com/xm-chentl/go-grafana/datasourcetype"
	"github.com/xm-chentl/go-grafana/targetformat"
	"github.com/xm-chentl/go-grafana/tpl"
)

type dashboardPanelBase struct {
	id         int32
	title      string
	YAxes      *DashboardPanelYAxes
	targets    []IDashboardPanelTarget
	thresholds []DashboardPanelThreshold
	overrides  []DashboardPanelOverride

	dbType    datasourcetype.Value
	chartType charttype.Value
}

func (d *dashboardPanelBase) SetInfo(id int32, title string) {
	d.id = id
	d.title = title
}

func (d *dashboardPanelBase) SetYAxes(yaxes *DashboardPanelYAxes) {
	d.YAxes = yaxes
}

func (d *dashboardPanelBase) AddTarget(target IDashboardPanelTarget) {
	d.targets = append(d.targets, target)
}

func (d *dashboardPanelBase) AddThreshold(thresholds ...DashboardPanelThreshold) {
	d.thresholds = append(d.thresholds, thresholds...)
}

func (d *dashboardPanelBase) AddOverride(overrides ...DashboardPanelOverride) {
	d.overrides = append(d.overrides, overrides...)
}

func (d dashboardPanelBase) Generate() (res string, err error) {
	if d.id == 0 {
		err = errors.New(ErrorPleaseSetPanelID)
		return
	}
	if d.title == "" {
		err = errors.New(ErrorPleaseSetPanelTitle)
		return
	}

	panelYAxes := "[]"
	if d.YAxes != nil {
		if panelYAxes, err = d.YAxes.Generate(); err != nil {
			return
		}
	}

	targetArray := make([]string, 0)
	if len(d.targets) > 0 {
		var targetStr string
		for _, item := range d.targets {
			if d.chartType == charttype.Table {
				item.SetFormat(targetformat.Table)
			}
			targetStr, err = item.Generate()
			if err != nil {
				return
			}
			targetArray = append(targetArray, targetStr)
		}
	} else {
		// 使用默认模板
		targetArray = append(targetArray, tpl.EmptyTarget)
	}

	thresholdArray := make([]string, 0)
	if len(d.thresholds) > 0 {
		var thresholdStr string
		for _, item := range d.thresholds {
			if thresholdStr, err = item.Generate(); err != nil {
				return
			}
			thresholdArray = append(thresholdArray, thresholdStr)
		}
	}

	overrides := make([]string, 0)
	if len(d.overrides) > 0 {
		var overrideItem string
		for _, item := range d.overrides {
			if overrideItem, err = item.Generate(); err != nil {
				return
			}
			overrides = append(overrides, overrideItem)
		}
	}

	var panelTemplateContent string
	switch d.chartType {
	case charttype.Histogram, charttype.Stitches:
		panelTemplateContent = tpl.PanelStitchesOrHistogram
	case charttype.Pie:
		panelTemplateContent = tpl.PanelPie
	case charttype.Star:
		panelTemplateContent = tpl.PanelStar
	case charttype.Table:
		panelTemplateContent = tpl.PanelTable
	}

	var panelTpl *template.Template
	panelTpl, err = template.New("dashboard-panel-default").Funcs(tplFuncs).Parse(panelTemplateContent)
	if err != nil {
		return
	}

	// 拆线和柱状与其它图表有差异
	tplData := map[string]interface{}{
		"id":         d.id,
		"title":      d.title,
		"targets":    targetArray,
		"datasource": "",
	}

	switch d.chartType {
	case charttype.Histogram:
		tplData["bars"] = true
		tplData["lines"] = false
		tplData["thresholds"] = thresholdArray
		tplData["overrides"] = overrides
		tplData["yaxes"] = panelYAxes
	case charttype.Stitches:
		tplData["bars"] = false
		tplData["lines"] = true
		tplData["thresholds"] = thresholdArray
		tplData["overrides"] = overrides
		tplData["yaxes"] = panelYAxes
	case charttype.Star, charttype.Table:
		tplData["overrides"] = overrides
	}

	var bf bytes.Buffer
	err = panelTpl.Execute(&bf, tplData)
	if err != nil {
		return
	}
	res = bf.String()

	return
}
