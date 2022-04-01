package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1608 lang=golang
 *
 * [1608] 特殊数组的特征值
 */

// @lc code=start
func specialArray(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		special := len(nums) - i
		if nums[i] >= special {
			if i <= 0 || nums[i-1] < special {
				return special
			}
		}
	}
	return -1
}

// @lc code=end

/*
	排序，然后判断
*/
