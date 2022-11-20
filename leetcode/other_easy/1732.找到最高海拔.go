package leetcode

/*
 * @lc app=leetcode.cn id=1732 lang=golang
 *
 * [1732] 找到最高海拔
 */

// @lc code=start
func largestAltitude(gain []int) int {
	max := 0
	current := 0
	for _, g := range gain {
		current += g
		if current > max {
			max = current
		}
	}
	return max
}

// @lc code=end

/*
	直接算出每个海拔点，然后判断即可
*/
