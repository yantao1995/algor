package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=1556 lang=golang
 *
 * [1556] 千位分隔数
 */

// @lc code=start
func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		temp := strconv.Itoa(n % 1000)
		for len(temp) < 3 {
			temp = "0" + temp
		}
		result = "." + temp + result
		n /= 1000
	}
	index := 1
	for {
		if result[index] != '0' {
			return result[index:]
		}
		index++
	}
}

// @lc code=end
