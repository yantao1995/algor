package leetcode

/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
	if n%2 != 0 && n != 1 {
		return false
	}
	i := 1
	for i <= n {
		if i == n {
			return true
		}
		i *= 2
	}
	return false
}

// @lc code=end
