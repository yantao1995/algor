package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=868 lang=golang
 *
 * [868] 二进制间距
 */

// @lc code=start
func binaryGap(n int) int {
	str := ""
	for n > 0 {
		str += strconv.Itoa(n % 2)
		n /= 2
	}
	for k := range str {
		if str[k] == '1' {
			str = str[k:]
			break
		}
	}
	if len(str) == 1 {
		return 0
	}
	slice := strings.Split(str, "1")
	count := 0
	for k := range slice {
		if len(slice[k]) > count {
			count = len(slice[k])
		}
	}
	return 1 + count
}

// @lc code=end
