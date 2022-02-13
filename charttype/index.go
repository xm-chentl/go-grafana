package charttype

// Value 图表枚举
type Value int32

const (
	// Stitches 拆线图
	Stitches Value = iota + 1
	// Histogram 柱状图
	Histogram
	// Pie 饼图
	Pie
	// Star 指标卡
	Star
	// Table 表格
	Table
)
