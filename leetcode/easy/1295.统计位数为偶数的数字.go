package easy

import "strconv"

/*
 * @lc app=leetcode.cn id=1295 lang=golang
 *
 * [1295] 统计位数为偶数的数字
 */

// @lc code=start
func findNumbers(nums []int) int {
	result := 0
	for k := range nums {
		if len(strconv.Itoa(nums[k]))%2 == 0 {
			result++
		}
	}
	return result
}

// @lc code=end
