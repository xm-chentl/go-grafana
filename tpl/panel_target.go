package tpl

var (

	// TargetTpl 多数据源请求模板
	// .hide 是否隐藏
	// .alias 别名
	// .dsType 数据库类型
	// .table 数据库表名(metric)
	// .query 查询语句
	// .refID 查询别名
	// .resultFormat 结果格式
	// .datasource 数据库（Grafana配置的数据源）
	Target = `{
		"hide":{{.hide}},
		"alias": "{{.alias}}",
		"dsType": "{{.dsType}}",
		"datasource": "{{.datasource}}",
		"groupBy": [{{ join .groupBy "," }}],
		"measurement": "{{.table}}",
		"orderByTime": "ASC",
		"policy": "default",
		"query": "{{.query}}",
		"rawQuery": true,
		"refId": "{{.refID}}",
		"resultFormat": "{{.format}}",
		"select": [{{ join .select " " }}],
		"tags": []
	}`

	// TargetVictoria 维多利亚Target模板
	// .hide 是否隐藏
	// .query 查询
	// .format 格式
	// .interval 最小步频(min step)
	// .legendFormat 对应legend
	// .refID 查询别名
	// .datasource 数据库 (Grafana配置的数据源)
	TargetVictoria = `{
		"datasource": "{{.datasource}}",
		"hide": {{.hide}},
		"exemplar": true,
		"expr": "{{.query}}",
		"format": "{{.format}}",
		"interval": "{{.interval}}",
		"legendFormat": "{{.legendFormat}}",
		"refId": "{{.refID}}"
	  }`

	// TplEmptyTarget 默认target模板
	EmptyTarget = `{
		"groupBy": [
		  {
			"params": [
			  "$__interval"
			],
			"type": "time"
		  },
		  {
			"params": [
			  "null"
			],
			"type": "fill"
		  }
		],
		"hide":false,
		"orderByTime": "ASC",
		"policy": "default",
		"refId": "A",
		"resultFormat": "time_series",
		"select": [
		  [
			{
			  "params": [
				"value"
			  ],
			  "type": "field"
			},
			{
			  "params": [],
			  "type": "mean"
			}
		  ]
		],
		"tags": []
	  }`

	// TplParam 参数模板（用于Target结构下 groupBy、select）
	// .params 值(这个是多个，未处理)
	// .type 类型或者函数名
	Param = `{
		"params": [
		  "{{.params}}"
		],
		"type": "{{.type}}"
	  }`
)
