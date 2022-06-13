package leetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=504 lang=golang
 *
 * [504] 七进制数
 */

// @lc code=start
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	symbol := ""
	if num < 0 {
		symbol = "-"
		num = 0 - num
	}
	result := ""
	for num > 0 {
		result = strconv.Itoa(num%7) + result
		num /= 7
	}
	result = symbol + result
	return result
}

// @lc code=end
