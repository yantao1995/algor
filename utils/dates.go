package utils

import (
	"time"
)

const SysTimeform = "2006-01-02 15:04:05"

const SysTimeformShort = "2006-01-02"

var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing") // 系统时区

//小时开始结束时间
func HourStartAndEndTime(unix int) (string, int, int) {
	t := time.Unix(int64(unix), 0).In(SysTimeLocation)
	t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, SysTimeLocation)
	dayStartDate, dayStartUnix := t.Format("2006-01-02"), int(t.Unix())
	return dayStartDate, dayStartUnix, dayStartUnix + 60*60 - 1
}

//天开始结束时间
func DayStartAndEndTime(unix int) (string, int, int) {
	t := time.Unix(int64(unix), 0).In(SysTimeLocation)
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, SysTimeLocation)
	dayStartDate, dayStartUnix := t.Format("2006-01-02"), int(t.Unix())
	return dayStartDate, dayStartUnix, dayStartUnix + 24*60*60 - 1
}

//周开始结束时间
func WeekStartAndEndTime(unix int) (string, string, int, int) {
	t := time.Unix(int64(unix), 0).In(SysTimeLocation)
	mondayDate, mondayUnix := getFirstDateOfWeek(t)
	sundayDate, sundayUnix := getLastDateOfWeek(t)
	return mondayDate, sundayDate, mondayUnix, sundayUnix
}

// 获取传入时间的周一 的日期
func getFirstDateOfWeek(t time.Time) (string, int) {
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStartDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, SysTimeLocation).AddDate(0, 0, offset)
	return weekStartDate.Format("2006-01-02"), int(weekStartDate.Unix())
}

// 获取传入时间的周日 的日期
func getLastDateOfWeek(t time.Time) (string, int) {
	offset := int(t.Weekday() - time.Sunday)
	if offset > 0 {
		offset = 7 - offset
	}
	weekEndDate := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, SysTimeLocation).AddDate(0, 0, offset)
	return weekEndDate.Format("2006-01-02"), int(weekEndDate.Unix())
}

//月开始结束时间
func MonthStartAndEndTime(unix int) (string, string, int, int) {
	t := time.Unix(int64(unix), 0).In(SysTimeLocation)
	startDate := t.AddDate(0, 0, -t.Day()+1).Format("2006-01-02")
	endDate := t.AddDate(0, 0, -t.Day()+1).AddDate(0, 1, -1).Format("2006-01-02")
	startUnix, _ := time.ParseInLocation("2006-01-02 15:04:05", startDate+" 00:00:00", SysTimeLocation)
	endUnix, _ := time.ParseInLocation("2006-01-02 15:04:05", endDate+" 23:59:59", SysTimeLocation)
	return startDate, endDate, int(startUnix.Unix()), int(endUnix.Unix())
}

// 获取时间所在月的开始和结束日期 返回列表(开始时间戳，结束时间戳，简短开始时间，简短结束时间，详细开始时间，详细结束时间)
func GetMonthStartAndEndTimeUnixStrs(timeInt int64) (int64, int64, string, string, string, string) {
	if timeInt <= 0 {
		return 0, 0, "", "", "", ""
	}
	t := time.Unix(timeInt, 0)
	startDate := t.AddDate(0, 0, -t.Day()+1).Format("2006-01-02")
	endDate := t.AddDate(0, 0, -t.Day()+1).AddDate(0, 1, -1).Format("2006-01-02")
	startDateStr, endDateStr := startDate+" 00:00:00", endDate+" 23:59:59"
	startUnix, _ := time.ParseInLocation(SysTimeform, startDateStr, SysTimeLocation)
	endUnix, _ := time.ParseInLocation(SysTimeform, endDateStr, SysTimeLocation)
	return startUnix.Unix(), endUnix.Unix(), startDate, endDate, startDateStr, endDateStr
}

// 获取时间所在年的开始和结束日期 返回列表(开始时间戳，结束时间戳，简短开始时间，简短结束时间，详细开始时间，详细结束时间)
func GetYearStartAndEndTimeUnixStrs(timeInt int64) (int64, int64, string, string, string, string) {
	if timeInt <= 0 {
		return 0, 0, "", "", "", ""
	}
	t := time.Unix(timeInt, 0)
	startDate := t.AddDate(0, -int(t.Month())+1, -t.Day()+1).Format("2006-01-02")
	endDate := t.AddDate(1, -int(t.Month())+1, -t.Day()).Format("2006-01-02")
	startDateStr, endDateStr := startDate+" 00:00:00", endDate+" 23:59:59"
	startUnix, _ := time.ParseInLocation(SysTimeform, startDateStr, SysTimeLocation)
	endUnix, _ := time.ParseInLocation(SysTimeform, endDateStr, SysTimeLocation)
	return startUnix.Unix(), endUnix.Unix(), startDate, endDate, startDateStr, endDateStr
}
