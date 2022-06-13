package leetcode

/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 缺失数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	bit := make([]int, len(nums)+1)
	for k := range nums {
		bit[nums[k]] = 1
	}
	for k := range bit {
		if bit[k] == 0 {
			return k
		}
	}
	return 0
}

// @lc code=end
