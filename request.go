package gografana

import "time"

type RequestPreviewDashboard struct {
	PanelID  int32
	Url      string
	FromTime *time.Time
	ToTime   *time.Time
	Vars     map[string][]string
}
