package easy

import "sort"

/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] 三个数的最大乘积
 */

// @lc code=start
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	r1 := nums[0] * nums[1] * nums[len(nums)-1]
	r2 := nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
	if r1 > r2 {
		return r1
	}
	return r2
}

// @lc code=end
