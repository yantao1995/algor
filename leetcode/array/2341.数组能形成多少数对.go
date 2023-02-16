package leetcode

/*
 * @lc app=leetcode.cn id=2341 lang=golang
 *
 * [2341] 数组能形成多少数对
 */

// @lc code=start
func numberOfPairs(nums []int) []int {
	result := []int{0, len(nums)}
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] > 0 && m[nums[i]]&1 == 0 {
			result[0]++
			result[1] -= 2
		}
	}
	return result
}

// @lc code=end
