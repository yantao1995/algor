package utils

import (
	"time"
)

// GetZeroUnix 获取指定时区、时间戳的0点时间戳
func GetZeroUnix(loc *time.Location, ts int64) int64 {
	t := time.Unix(ts, 0).In(loc)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc).Unix()
}

// GetTodayZeroUnix 获取指定时区当天0点时间戳
func GetTodayZeroUnix(loc *time.Location) int64 {
	return GetZeroUnix(loc, time.Now().Unix())
}

// 获取月开始时间
func GetMonthStart(loc *time.Location, ts int64) int64 {
	t := time.Unix(ts, 0).In(loc)
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, loc).Unix()
}

// 获取月结束时间
func GetMonthEnd(loc *time.Location, ts int64) int64 {
	t := time.Unix(ts, 0).In(loc)
	return time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, loc).Unix() - 1
}
