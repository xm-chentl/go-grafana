package targetformat

type Value string

func (v Value) String() string {
	return string(v)
}

const (
	Empty      Value = ""
	TimeSeries Value = "time_series"
	Table      Value = "table"
	Heatmap    Value = "heatmap"
)
