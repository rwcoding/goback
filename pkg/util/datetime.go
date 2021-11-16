package util

import (
	"time"
)

// StringToUnix 字符串转换为unix时间戳，单位为秒
func StringToUnix(dt string) (uint32, error) {
	l := len(dt)
	timeFormat := "2006-01-02 15:04:05"
	if l == 10 {
		timeFormat = "2006-01-02"
	} else if l == 8 {
		timeFormat = "15:04:05"
	}
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return 0, err
	}
	t, err := time.ParseInLocation(timeFormat, dt, loc)
	if err != nil {
		return 0, err
	}
	return uint32(t.Unix()), nil
}
