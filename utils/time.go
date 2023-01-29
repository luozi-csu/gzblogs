package utils

import "time"

const (
	stdDateStr string = "2006-01-02"
)

func Zerotime() int64 {
	timeStr := time.Now().Format(stdDateStr)
	t, _ := time.ParseInLocation(stdDateStr, timeStr, time.Local)
	return t.Unix()
}
