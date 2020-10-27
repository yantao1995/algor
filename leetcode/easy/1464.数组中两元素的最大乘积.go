package easy

/*
 * @lc app=leetcode.cn id=1464 lang=golang
 *
 * [1464] 数组中两元素的最大乘积
 */

// @lc code=start
func maxProduct(nums []int) int {
	max := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i]-1)*(nums[j]-1) > max {
				max = (nums[i] - 1) * (nums[j] - 1)
			}
		}
	}
	return max
}

// @lc code=end
