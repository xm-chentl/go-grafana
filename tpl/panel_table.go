package tpl

var (
	// PanelTable 表格图表
	// .id 图表id
	// .title 图表标题
	// .targets 查询
	// .overrides 别名
	PanelTable = `{
		"datasource": "-- Mixed --",
		"description": "",
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "thresholds"
			},
			"custom": {
			  "align": null,
			  "filterable": false
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green",
				  "value": null
				},
				{
				  "color": "red",
				  "value": 80
				}
			  ]
			}
		  },
		  "overrides": [{{ join .overrides "," }}]
		},
		"gridPos": {
		  "h": 8,
		  "w": 12,
		  "x": 0,
		  "y": 0
		},
		"id": {{.id}},
		"options": {
		  "showHeader": true,
		  "sortBy": [
			{
			  "desc": true,
			  "displayName": "core95"
			}
		  ]
		},
		"pluginVersion": "7.5.8",
		"targets": [{{ join .targets "," }}],
		"timeFrom": null,
		"timeShift": null,
		"title": "{{.title}}",
		"type": "table"
	  }`
)
