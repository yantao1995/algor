package easy

/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3的幂
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	if n == 1 {
		return true
	}
	return false
}

// @lc code=end
