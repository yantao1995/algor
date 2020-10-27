package easy

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=1317 lang=golang
 *
 * [1317] 将整数转换为两个无零整数的和
 */

// @lc code=start
func getNoZeroIntegers(n int) []int {
	for i := 1; i <= n/2; i++ {
		if !strings.Contains(strconv.Itoa(i), "0") &&
			!strings.Contains(strconv.Itoa(n-i), "0") {
			return []int{i, n - i}
		}
	}
	return nil
}

// @lc code=end
