package models

type ComResult struct {
	Code     string
	Msg      string
	TimeUnix int64
	Data     interface{}
}
