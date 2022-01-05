package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=164 lang=golang
 *
 * [164] 最大间距
 */

// @lc code=start
func maximumGap(nums []int) int {
	sort.Ints(nums)
	max := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > max {
			max = nums[i] - nums[i-1]
		}
	}
	return max
}

// @lc code=end
