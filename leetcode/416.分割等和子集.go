package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	sort.Ints(nums)
	prev, tail := 0, nums[len(nums)-1]
	for i := 0; i < len(nums)-1; i++ {
		prev += nums[i]
	}
	if prev == tail {
		return true
	}
	for j := len(nums) - 2; j > 0 && prev > tail; j-- {
		prev -= nums[j]
		tail += nums[j]
		if prev == tail {
			return true
		}
	}
	return false
}

// @lc code=end
