package utils

import (
	"time"
)

const SysTimeform = "2006-01-02 15:04:05"

const SysTimeformShort = "2006-01-02"

var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing") // 系统时区

// 获取时间所在周日日期
func GetLastDateOfWeek(timeInt int64) int {
	if timeInt == 0 {
		return 0
	}
	t := time.Unix(timeInt, 0)
	offset := int(t.Weekday() - time.Sunday)
	if offset > 0 {
		offset = 7 - offset
	}
	weekStartDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekSunday := weekStartDate.Format("2006-01-02") + " 23:59:59"
	weekMondayUnix, _ := time.ParseInLocation(SysTimeform, weekSunday, SysTimeLocation)

	return int(weekMondayUnix.Unix())
}

// 获取时间所在月的开始和结束日期 返回列表(开始时间戳，结束时间戳，简短开始时间，简短结束时间，详细开始时间，详细结束时间)
func GetMonthStartAndEndTimeUnixStrs(timeInt int64) (int64, int64, string, string, string, string) {
	if timeInt <= 0 {
		return 0, 0, "", "", "", ""
	}
	t := time.Unix(timeInt, 0)
	startDate := t.AddDate(0, 0, -t.Day()+1).Format("2006-01-02")
	endDate := t.AddDate(0, 0, -t.Day()+1).AddDate(0, 1, -1).Format("2006-01-02")
	yestdayStartStr, yestdayEndStr := startDate+" 00:00:00", endDate+" 23:59:59"
	startUnix, _ := time.ParseInLocation(SysTimeform, yestdayStartStr, SysTimeLocation)
	endUnix, _ := time.ParseInLocation(SysTimeform, yestdayEndStr, SysTimeLocation)
	return startUnix.Unix(), endUnix.Unix(), startDate, endDate, yestdayStartStr, yestdayEndStr
}

// 获取时间所在年的开始和结束日期 返回列表(开始时间戳，结束时间戳，简短开始时间，简短结束时间，详细开始时间，详细结束时间)
func GetYearStartAndEndTimeUnixStrs(timeInt int64) (int64, int64, string, string, string, string) {
	if timeInt <= 0 {
		return 0, 0, "", "", "", ""
	}
	t := time.Unix(timeInt, 0)
	startDate := t.AddDate(0, -int(t.Month())+1, -t.Day()+1).Format("2006-01-02")
	endDate := t.AddDate(1, -int(t.Month())+1, -t.Day()).Format("2006-01-02")
	yestdayStartStr, yestdayEndStr := startDate+" 00:00:00", endDate+" 23:59:59"
	startUnix, _ := time.ParseInLocation(SysTimeform, yestdayStartStr, SysTimeLocation)
	endUnix, _ := time.ParseInLocation(SysTimeform, yestdayEndStr, SysTimeLocation)
	return startUnix.Unix(), endUnix.Unix(), startDate, endDate, yestdayStartStr, yestdayEndStr
}
