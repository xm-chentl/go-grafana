package tpl

var (
	// PanelPie 饼图模板
	// .id 图表id
	// .title 图表标题
	// .targets 查询
	PanelPie = `{
		"aliasColors": {},
		"breakPoint": "50%",
		"cacheTimeout": null,
		"combine": {
		  "label": "Others",
		  "threshold": 0
		},
		"datasource": "-- Mixed --",
		"fieldConfig": {
		  "defaults": {},
		  "overrides": []
		},
		"fontSize": "80%",
		"format": "short",
		"gridPos": {
		  "h": 8,
		  "w": 12,
		  "x": 0,
		  "y": 0
		},
		"id": {{.id}},
		"interval": null,
		"legend": {
		  "show": true,
		  "values": true
		},
		"legendType": "Under graph",
		"links": [],
		"maxDataPoints": 1,
		"nullPointMode": "connected",
		"pieType": "pie",
		"pluginVersion": "7.5.8",
		"strokeWidth": 1,
		"targets": [{{ join .targets "," }}],
		"title": "{{.title}}",
		"type": "grafana-piechart-panel",
		"valueName": "current"
	  }`
)
