package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=2357 lang=golang
 *
 * [2357] 使数组中所有元素都等于零
 */

// @lc code=start
func minimumOperations(nums []int) int {
	sort.Ints(nums)
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			for j := i + 1; j < len(nums); j++ {
				nums[j] -= nums[i]
			}
			result++
			nums[i] = 0
		}
	}
	return result
}

// @lc code=end

/*
	排序后直接依次减即可
*/
