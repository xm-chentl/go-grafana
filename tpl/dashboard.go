package tpl

const (
	// DashboardTpl 请求模板
	// .id 标识（用于修改）
	// .title 标题
	// .panels 图表集
	// .refresh 刷新间隔
	// .folderID 文件夹标识
	// .message 仪表盘说明
	// .version 版本，当修改了版本之后，这个需要做相应的累加
	// .templatingList 仪表盘变量
	Dashboard = `{
		"dashboard": {
		  "id": {{.id}},
		  "uid": null,
		  "title": "{{.title}}",
		  "tags": [],
		  "panels": [{{ join .panels ","}}],
		  "timezone": "browser",
		  "schemaVersion": 16,
		  "version": {{.version}},
		  "refresh": "{{.refresh}}",
		  "templating": {
			"list": [{{ join .templatingList "," }}]
		  }
		},
		"folderId": {{.folderID}},
		"folderUid": "",
		"message": "{{.message}}",
		"overwrite": false
	}`
	// DashboardTemplatingListItem 变量选集
	// .datasource 数据源(数据库)
	// .query 查询语句 如：SHOW tag values from \"re__node_cache_map\" WITH key = \"node\"  WHERE \"cache\"='$cache'
	// .name 变量名
	// .label 标签名
	// .multi 是否多选
	// .desc 描述
	// .type 变量模式
	// .options 选项值
	// .current 当前选中
	// .all 是否加入all选项
	DashboardTemplatingListItem = `{
        "allValue": null,
        "current": {{.current}},
        "datasource": "{{.datasource}}",
        "definition": "{{.query}}",
        "description": "{{.desc}}",
        "error": null,
        "hide": 0,
        "includeAll": {{.all}},
        "label": "{{.label}}",
        "multi": {{.multi}},
        "name": "{{.name}}",
        "options": [{{ join .options "," }}],
        "query": "{{.query}}",
        "refresh": 1,
        "regex": "",
        "skipUrlSync": false,
        "sort": 0,
        "tagValuesQuery": "",
        "tags": [],
        "tagsQuery": "",
        "type": "{{.type}}",
        "useTags": false
      }`
	// DashboardTemplatingOptionItem 变量选择项
	// .selected 是否选中
	// .text 选项名
	// .value 选项值
	DashboardTemplatingOptionItem = `{
		"selected": {{.selected}},
		"text": "{{.text}}",
		"value": "{{.value}}"
	  }`
)
