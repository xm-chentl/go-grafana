package gografana

type ResponseErrorItem struct {
	FieldNames     []string `json:"fieldNames"`
	Classification string   `json:"classification"`
	Message        string   `json:"message"`
}

type ResponseBase struct {
	Code    ErrorCode
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseCreateOrUpdateDashboard struct {
	ResponseBase

	ID      int32  `json:"id"`
	UID     string `json:"uid"`
	Url     string `json:"url"`
	Version int32  `json:"version"`
}

type ResponseCreateOrUpdateFolder struct {
	ResponseBase

	ID          int32  `json:"id"`
	UID         string `json:"uid"`
	Name        string `json:"title"`
	Url         string `json:"url"`
	CanSave     bool   `json:"canSave"`
	CanEdit     bool   `json:"canEdit"`
	CanAdmin    bool   `json:"canAdmin"`
	CreatorName string `json:"createdBy"`
	CreatedAt   string `json:"created"`
	UpdatorName string `json:"updatedBy"`
	UpdatedAt   string `json:"updated"`
	Version     int32  `json:"version"`
}

type ResponseDeleteFolder struct {
	ResponseBase

	Message string `json:"message"`
	ID      int    `json:"id"`
}
