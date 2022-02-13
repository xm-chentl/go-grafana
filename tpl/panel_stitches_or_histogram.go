package tpl

var (
	// PanelStitchesOfHistogramTpl 拆线或柱状图请求参数模板
	// .bars 是否柱状(true、false)
	// .lines 是否拆线图(true、false)
	// .datasource 数据源数库(如: influxdb_os_backup)
	// .title 图表名
	// .scopedVars 变量值(暂时不支持)
	// .overrides 别名
	// .yaxes 单位配置
	PanelStitchesOrHistogram = `{
		"aliasColors": {},
		"bars": {{.bars}},
		"dashLength": 10,
		"dashes": false,
		"datasource": "-- Mixed --",
		"editable": true,
		"error": false,
		"fieldConfig": {
		  "defaults": {},
		  "overrides": [{{ join .overrides ","}}]
		},
		"fill": 4,
		"fillGradient": 0,
		"grid": {},
		"gridPos": {
		  "h": 8,
		  "w": 12,
		  "x": 0,
		  "y": 0
		},
		"height": "",
		"hiddenSeries": false,
		"id": {{.id}},
		"legend": {
		  "avg": false,
		  "current": false,
		  "max": false,
		  "min": false,
		  "show": true,
		  "total": false,
		  "values": false
		},
		"lines": {{.lines}},
		"linewidth": 1,
		"links": [],
		"nullPointMode": "null",
		"options": {
		  "alertThreshold": true
		},
		"paceLength": 10,
		"percentage": false,
		"pluginVersion": "7.5.8",
		"pointradius": 1,
		"points": false,
		"renderer": "flot",
		"scopedVars": {},
		"seriesOverrides": [],
		"spaceLength": 10,
		"stack": false,
		"steppedLine": false,
		"targets": [{{ join .targets "," }}],
		"thresholds": [{{ join .thresholds "," }}],
		"timeFrom": null,
		"timeRegions": [],
		"timeShift": null,
		"title": "{{.title}}",
		"tooltip": {
		  "shared": true,
		  "sort": 0,
		  "value_type": "cumulative"
		},
		"type": "graph",
		"xaxis": {
		  "buckets": null,
		  "mode": "time",
		  "name": null,
		  "show": true,
		  "values": []
		},
		"yaxes": {{.yaxes}},
		"yaxis": {
		  "align": false,
		  "alignLevel": null
		}
	}`
)
