package tpl

var (
	// PanelStar 指标卡
	// .id 图表id
	// .title 图表标题
	// .targets 查询
	// .overrides 别名配置
	PanelStar = `{
		"datasource": "-- Mixed --",
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "thresholds"
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
		  "overrides": [{{ join .overrides ","}}]
		},
		"gridPos": {
		  "h": 8,
		  "w": 12,
		  "x": 0,
		  "y": 0
		},
		"id": {{.id}},
		"options": {
		  "colorMode": "value",
		  "graphMode": "area",
		  "justifyMode": "auto",
		  "orientation": "auto",
		  "reduceOptions": {
			"calcs": [
			  "lastNotNull"
			],
			"fields": "",
			"values": false
		  },
		  "text": {},
		  "textMode": "auto"
		},
		"pluginVersion": "7.5.8",
		"targets": [{{ join .targets "," }}],
		"timeFrom": null,
		"timeShift": null,
		"title": "{{.title}}",
		"type": "stat"
	  }`
)
