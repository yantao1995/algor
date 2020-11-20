package leetcode

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {

	m = map[int]int{}
	return dp198(nums, 0)
}

var m map[int]int = map[int]int{}

func dp198(nums []int, index int) int {
	if index >= len(nums) {
		return 0
	}
	if v, ok := m[index]; ok {
		return v
	}
	m[index] = max198(dp198(nums, index+1),
		nums[index]+dp198(nums, index+2))
	return m[index]
}

func max198(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end
