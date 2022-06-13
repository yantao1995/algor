package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=747 lang=golang
 *
 * [747] 至少是其他数字两倍的最大数
 */

// @lc code=start
func dominantIndex(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	if temp[len(temp)-1]-temp[len(temp)-2] >= temp[len(temp)-2] {
		for k := range nums {
			if temp[len(temp)-1] == nums[k] {
				return k
			}
		}
	}
	return -1
}

// @lc code=end
