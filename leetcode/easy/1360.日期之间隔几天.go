package leetcode

import "time"

/*
 * @lc app=leetcode.cn id=1360 lang=golang
 *
 * [1360] 日期之间隔几天
 */

// @lc code=start
func daysBetweenDates(date1 string, date2 string) int {
	t1, _ := time.Parse("2006-01-02", date1)
	t2, _ := time.Parse("2006-01-02", date2)
	if t2.Unix() > t1.Unix() {
		return int((t2.Unix() - t1.Unix()) / (24 * 60 * 60))
	}
	return int((t1.Unix() - t2.Unix()) / (24 * 60 * 60))
}

// @lc code=end
