package easy

/*
 * @lc app=leetcode.cn id=1486 lang=golang
 *
 * [1486] 数组异或操作
 */

// @lc code=start
func xorOperation(n int, start int) int {
	result := start
	for n > 1 {
		n--
		result ^= (start + 2*n)
	}
	return result
}

// @lc code=end
