package leetcode

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		current := 1
		for j := i; j < len(nums); j++ {
			current *= nums[j]
			if current > max {
				max = current
			}
			if nums[j] == 0 {
				break
			}
		}
	}
	return max
}

// @lc code=end
