package leetcode

import (
	"strconv"
	"time"
)

/*
 * @lc app=leetcode.cn id=1185 lang=golang
 *
 * [1185] 一周中的第几天
 */

// @lc code=start
func dayOfTheWeek(day int, month int, year int) string {
	//1970-01-01 星期四
	weekUnix := 7 * 24 * 60 * 60
	tStr := strconv.Itoa(year) + "-" + strconv.Itoa(month) + "-" + strconv.Itoa(day)
	t, _ := time.Parse("2006-1-2", tStr)
	count := (int(t.Unix()) + 4*24*60*60) % weekUnix
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	return days[count/(24*60*60)]
}

// @lc code=end
