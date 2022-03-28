package leetcode

/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	for k := range result {
		result[k] = 1
	}
	for k := range nums {
		for kk := range result {
			if k != kk {
				result[kk] = result[kk] * nums[k]
			}
		}
	}
	return result
}

// @lc code=end
