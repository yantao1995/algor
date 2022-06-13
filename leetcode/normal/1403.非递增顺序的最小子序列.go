package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1403 lang=golang
 *
 * [1403] 非递增顺序的最小子序列
 */

// @lc code=start
func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	sum := 0
	for k := range nums {
		sum += nums[k]
	}
	half := 0
	result := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		result = append(result, nums[i])
		if half+nums[i] > sum/2 {
			return result
		}
		half += nums[i]
	}
	return result
}

// @lc code=end
