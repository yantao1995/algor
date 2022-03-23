package leetcode

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	max := -1 << 31
	current := 0
	for i := 0; i < len(nums); i++ {
		current += nums[i]
		if max < current {
			max = current
		}
		if current < 0 {
			current = 0
		}
	}
	return max
}

// @lc code=end
