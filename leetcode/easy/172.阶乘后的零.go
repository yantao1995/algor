package easy

/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */

// @lc code=start
func trailingZeroes(n int) int {
	val := 0
	for n >= 5 {
		n /= 5
		val += n
	}
	return val
}

// @lc code=end
