package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	index := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if index == k {
			return nums[i]
		}
		index++
	}
	return 0
}

// @lc code=end
