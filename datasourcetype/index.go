package datasourcetype

type Value int

const (
	MySql Value = iota + 1
	Influxdb
	Victoria
)
