package leetcode

/*
 * @lc app=leetcode.cn id=1304 lang=golang
 *
 * [1304] 和为零的N个唯一整数
 */

// @lc code=start
func sumZero(n int) []int {
	result := []int{}
	temp := n / 2
	for temp > 0 {
		result = append(result, temp)
		result = append(result, 0-temp)
		temp--
	}
	if n%2 == 1 {
		result = append(result, 0)
	}
	return result
}

// @lc code=end
