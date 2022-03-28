package leetcode

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	max := nums[0]
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {

		for j := i; j < len(nums); j++ {

		}
	}
	return max
}

// @lc code=end
