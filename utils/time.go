package utils

import "time"

const (
	timeFormat string = "2006-01-02"
)

func Zerotime() int64 {
	timeStr := time.Now().Format(timeFormat)
	t, _ := time.ParseInLocation(timeFormat, timeStr, time.Local)
	return t.Unix()
}
