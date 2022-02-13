package gografana

import (
	"bytes"
	"text/template"

	"github.com/xm-chentl/go-grafana/tpl"
)

type DashboardPanelOverride struct {
	Alias   string // 别名
	Options string // monitor_is_py_alive.mean
}

func (o DashboardPanelOverride) Generate() (res string, err error) {
	var overrideTpl *template.Template
	overrideTpl, err = template.New("dashboard-panel-overrides").Parse(tpl.PanelOverrideItem)
	if err != nil {
		return
	}

	var by bytes.Buffer
	err = overrideTpl.Execute(&by, map[string]interface{}{
		"options": o.Options,
		"value":   o.Alias,
	})
	if err != nil {
		return
	}
	res = by.String()

	return
}

type DashboardPanelYAxesConfig struct {
	Format   string
	Label    string
	Min      int
	Max      int
	Decimals int
}

type DashboardPanelYAxes struct {
	Left  DashboardPanelYAxesConfig
	Right DashboardPanelYAxesConfig
}

func (d DashboardPanelYAxes) Generate() (res string, err error) {
	var overrideTpl *template.Template
	overrideTpl, err = template.New("dashboard-panel-yaxes").Parse(tpl.PanelYaxes)
	if err != nil {
		return
	}

	var by bytes.Buffer
	err = overrideTpl.Execute(&by, map[string]interface{}{
		"leftFormat":    d.Left.Format,
		"leftLabel":     d.Left.Label,
		"leftMin":       d.Left.Min,
		"leftMax":       d.Left.Max,
		"leftDecimals":  d.Left.Decimals,
		"rightFormat":   d.Right.Format,
		"rightLabel":    d.Right.Label,
		"rightMin":      d.Right.Min,
		"rightMax":      d.Right.Max,
		"rightDecimals": d.Right.Decimals,
	})
	if err != nil {
		return
	}
	res = by.String()

	return
}
