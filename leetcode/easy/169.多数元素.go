package easy

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	m := map[int]int{}
	for k := range nums {
		m[nums[k]]++
	}
	for k := range m {
		if m[k] > len(nums)/2 {
			return k
		}
	}
	return 0
}

// @lc code=end
