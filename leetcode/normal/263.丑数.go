package leetcode

/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
func isUgly(num int) bool {
	for num > 1 {
		temp := num
		if num%5 == 0 {
			num /= 5
		}
		if num%3 == 0 {
			num /= 3
		}
		if num%2 == 0 {
			num /= 2
		}
		if temp == num {
			return false
		}
	}
	if num == 1 {
		return true
	}
	return false
}

// @lc code=end
