package gografana

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/xm-chentl/go-grafana/targetformat"
	"github.com/xm-chentl/go-grafana/tpl"
)

var (
	tplFuncs = template.FuncMap{
		"join": func(elems []string, sep string) string {
			return strings.Join(elems, sep)
		},
	}
)

type DashboardPanelThreshold struct {
	Color string
	Op    string
	Value int32
}

func (d DashboardPanelThreshold) Generate() (res string, err error) {
	var panelThresholdTpl *template.Template
	if panelThresholdTpl, err = template.New("dashboard-panel-thresholds").Funcs(tplFuncs).Parse(tpl.Threshold); err != nil {
		return
	}

	var bf bytes.Buffer
	err = panelThresholdTpl.Execute(&bf, map[string]interface{}{
		"color": d.Color,
		"op":    d.Op,
		"value": d.Value,
	})
	if err != nil {
		return
	}
	res = bf.String()

	return
}

type DashboardPanelTarget struct {
	DashboardPanelTargetBase

	DsType string
	Metric string // 映射至table
}

func (d DashboardPanelTarget) Generate() (res string, err error) {
	selects := make([]string, 0)
	groupBy := make([]string, 0)
	var panelTargetTpl *template.Template
	if panelTargetTpl, err = template.New("dashboard-panel-target").Funcs(tplFuncs).Parse(tpl.Target); err != nil {
		return
	}
	if d.TargetFormat == targetformat.Empty {
		d.TargetFormat = targetformat.TimeSeries
	}

	var bf bytes.Buffer
	err = panelTargetTpl.Execute(&bf, map[string]interface{}{
		"datasource": d.DataSource,
		"query":      d.Query,
		"refID":      d.RefID,
		"hide":       d.Hide,
		"format":     d.TargetFormat.String(),
		"alias":      d.Alias,
		"dsType":     d.DsType,
		"table":      d.Metric,
		"select":     selects,
		"groupBy":    groupBy,
	})
	if err != nil {
		return
	}
	res = bf.String()

	return
}

type DashboardPanelVictoriaTarget struct {
	DashboardPanelTargetBase

	interval string
}

func (d DashboardPanelVictoriaTarget) Generate() (res string, err error) {
	var panelTargetTpl *template.Template
	if panelTargetTpl, err = template.New("dashboard-panel-target").Funcs(tplFuncs).Parse(tpl.TargetVictoria); err != nil {
		return
	}
	if d.TargetFormat == targetformat.Empty {
		d.TargetFormat = targetformat.TimeSeries
	}

	var interval string
	if d.interval == "" {
		interval = "1m"
	}

	var bf bytes.Buffer
	err = panelTargetTpl.Execute(&bf, map[string]interface{}{
		"datasource":   d.DataSource,
		"query":        d.Query,
		"refID":        d.RefID,
		"hide":         d.Hide,
		"format":       d.TargetFormat.String(),
		"interval":     interval,
		"legendFormat": d.Alias, // 暂时以这个来
	})
	if err != nil {
		return
	}
	res = bf.String()

	return
}

type DashboardTemplatingOption struct {
	Selected bool
	Text     string
	Value    string
}

func (d DashboardTemplatingOption) Generate() (res string, err error) {
	var templatingListTpl *template.Template
	if templatingListTpl, err = template.New("dashboard-templating-option").Funcs(tplFuncs).Parse(tpl.DashboardTemplatingOptionItem); err != nil {
		return
	}

	var bf bytes.Buffer
	err = templatingListTpl.Execute(&bf, map[string]interface{}{
		"selected": d.Selected,
		"text":     d.Text,
		"value":    d.Value,
	})
	if err != nil {
		return
	}
	res = bf.String()

	return
}

type DashboardTemplatingListItem struct {
	All        bool
	Datasource string
	Query      string
	Name       string
	Label      string
	Multi      bool
	Desc       string
	Type       string
	Options    []DashboardTemplatingOption
}

func (d DashboardTemplatingListItem) Generate() (res string, err error) {
	current := "{}"
	var options []string
	if len(d.Options) > 0 {
		var optStr string
		for index := range d.Options {
			if index == 0 {
				d.Options[index].Selected = true
			}

			optStr, err = d.Options[index].Generate()
			if err != nil {
				return
			}
			options = append(options, optStr)
		}
		current = options[0]
	}

	var templatingListTpl *template.Template
	if templatingListTpl, err = template.New("dashboard-templating-list-item").Funcs(tplFuncs).Parse(tpl.DashboardTemplatingListItem); err != nil {
		return
	}

	var bf bytes.Buffer
	err = templatingListTpl.Execute(&bf, map[string]interface{}{
		"datasource": d.Datasource,
		"name":       d.Name,
		"label":      d.Label,
		"multi":      d.Multi,
		"all":        d.All,
		"query":      d.Query,
		"desc":       d.Desc,
		"type":       d.Type,
		"current":    current,
		"options":    options,
	})
	if err != nil {
		return
	}
	res = bf.String()

	return
}

type DashboardInfo struct {
	ID             int32
	Title          string
	Refresh        string
	FolderID       int32
	Message        string
	Version        int32
	TemplatingList []DashboardTemplatingListItem
}

func generateDashboardArgs(
	info DashboardInfo,
	panels ...IDashboardPanel,
) (res string, err error) {
	panelArray := make([]string, 0)
	if len(panels) > 0 {
		var panelStr string
		for _, p := range panels {
			panelStr, err = p.Generate()
			if err != nil {
				return
			}
			panelArray = append(panelArray, panelStr)
		}
	}

	templatingList := make([]string, 0)
	if len(info.TemplatingList) > 0 {
		var itemStr string
		for _, t := range info.TemplatingList {
			if itemStr, err = t.Generate(); err != nil {
				return
			}
			templatingList = append(templatingList, itemStr)
		}
	}

	var dashbooardTpl *template.Template
	if dashbooardTpl, err = template.New("dashboard").Funcs(tplFuncs).Parse(tpl.Dashboard); err != nil {
		return
	}

	var bf bytes.Buffer
	err = dashbooardTpl.Execute(&bf, map[string]interface{}{
		"id":             info.ID,
		"title":          info.Title,
		"refresh":        info.Refresh,
		"folderID":       info.FolderID,
		"message":        info.Message,
		"panels":         panelArray,
		"version":        info.Version,
		"templatingList": templatingList,
	})
	res = bf.String()

	return
}
