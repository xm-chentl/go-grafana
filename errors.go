package gografana

type ErrorCode int32

const (
	// Success 成功
	Success ErrorCode = 200
	// ErrorsInvaildJSONOrField 错误（无效的 json、丢失或无效的字段等）
	ErrorsInvaildJSONOrField ErrorCode = 400
	// Unauthorized 未授权
	Unauthorized ErrorCode = 401
	// AccessDenied 拒绝访问
	AccessDenied ErrorCode = 403
	// NotFound 未找到
	NotFound ErrorCode = 404
	// PreconditionFailed 先决条件失败
	PreconditionFailed = 408
)

const (
	ErrorPleaseSetPanelID    = "请设置图表PanelID"
	ErrorPleaseSetPanelTitle = "请设置图表Title"
	ErrorVersionMismatch     = "版本匹配不上，已在其它地方被修改"
	ErrorSameName            = "仪表盘同名"
)

type GrafanaError error
