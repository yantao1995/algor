package leetcode

/*
 * @lc app=leetcode.cn id=1523 lang=golang
 *
 * [1523] 在区间范围内统计奇数数目
 */

// @lc code=start
func countOdds(low int, high int) int {
	result := (high - low) / 2
	if high%2 == 1 || low%2 == 1 {
		result += 1
	}
	return result
}

// @lc code=end
