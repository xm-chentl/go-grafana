package tpl

var (
	// Threshold 参考线模板
	// .color 颜色
	// .op
	// .value 值
	Threshold = `{
		"$$hashKey": "object:46",
		"colorMode": "custom",
		"fill": true,
		"fillColor": "{{.color}}",
		"line": true,
		"lineColor": "{{.color}}",
		"op": "{{.op}}",
		"value":{{.value}},
		"yaxis": "left"
	  }`
)
