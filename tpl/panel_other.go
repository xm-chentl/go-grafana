package tpl

var (
	// PanelOverrideItem 别名
	// .options 指定字段 格式：monitor_is_py_alive.mean
	// .value 别名
	PanelOverrideItem = `{
		"matcher": {
		  "id": "byName",
		  "options": "{{.options}}"
		},
		"properties": [
		  {
			"id": "displayName",
			"value": "{{.value}}"
		  }
		]
	  }`
	// PanelYaxes 单位配置
	// .leftFormat 格式
	// .leftLabel 别名
	// .leftMax 最大值
	// .leftMin 最小值
	// .leftDecimals 小数位
	// .rightFormat 格式
	// .rightLabel 别名
	// .rightMax 最大值
	// .rightMin 最小值
	// .rightDecimals 小数位
	PanelYaxes = `[
		{
		  "format": "{{.leftFormat}}",
		  "label": "{{.leftLabel}}",
		  "logBase": 1,
		  "max": {{.leftMax}},
		  "min": {{.leftMin}},
		  "show": true,
		  "decimals":{{.leftDecimals}}
		},
		{
		  "format": "{{.rightFormat}}",
		  "label": "{{.rightLabel}}",
		  "logBase": 1,
		  "max": {{.rightMax}},
		  "min": {{.rightMin}},
		  "show": true,
		  "decimals":{{.rightDecimals}}
		}
	  ]`
)
