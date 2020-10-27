package easy

/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
func singleNumber(nums []int) int {
	m := map[int]int{}
	for k := range nums {
		m[nums[k]]++
	}
	for k := range m {
		if m[k] == 1 {
			return k
		}
	}
	return 0
}

// @lc code=end
