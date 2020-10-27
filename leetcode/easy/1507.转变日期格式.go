package easy

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=1507 lang=golang
 *
 * [1507] 转变日期格式
 */

// @lc code=start
func reformatDate(date string) string {
	slice := strings.Split(date, " ")
	day := slice[0][:len(slice[0])-2]
	monthMap := map[string]string{"Jan": "01", "Feb": "02", "Mar": "03", "Apr": "04", "May": "05",
		"Jun": "06", "Jul": "07", "Aug": "08", "Sep": "09", "Oct": "10", "Nov": "11", "Dec": "12"}
	if len(day) < 2 {
		day = "0" + day
	}
	return slice[2] + "-" + monthMap[slice[1]] + "-" + day
}

// @lc code=end
