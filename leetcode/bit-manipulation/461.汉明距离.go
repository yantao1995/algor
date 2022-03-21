package leetcode

/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	result := x ^ y
	count := 0
	for result > 0 {
		if result%2 == 1 {
			count++
		}
		result = result >> 1
	}
	return count
}

// @lc code=end
