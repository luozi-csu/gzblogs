package utils

import "time"

func ZeroTime() int64 {
	timeStr := time.Now().Format("2001-07-18")
	t, _ := time.ParseInLocation("2001-07-18", timeStr, time.Local)
	return t.Unix()
}
