package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=462 lang=golang
 *
 * [462] 最少移动次数使数组元素相等 II
 */

// @lc code=start
func minMoves2(nums []int) int {
	sort.Ints(nums)
	result := 0
	mid := len(nums) / 2
	absf := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	for k := range nums {
		result += absf(nums[k], nums[mid])
	}
	return result
}

// @lc code=end

/*
	移动次数最少，可以取一个中位数来作为标定值
	然后依次对每个数做绝对值运算即可
*/
