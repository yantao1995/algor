package easy

import "strconv"

/*
 * @lc app=leetcode.cn id=693 lang=golang
 *
 * [693] 交替位二进制数
 */

// @lc code=start
func hasAlternatingBits(n int) bool {
	if n == 0 {
		return true
	}
	result := ""
	for n > 0 {
		result = strconv.Itoa(n%2) + result
		n /= 2
	}
	for i := 0; i < len(result)-1; i++ {
		if result[i] == result[i+1] {
			return false
		}
	}
	return true
}

// @lc code=end
