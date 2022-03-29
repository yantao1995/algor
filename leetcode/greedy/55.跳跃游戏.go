package leetcode

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	max := 0
	for i := 0; i <= max && i < len(nums); i++ {
		if nums[i]+i > max {
			max = nums[i] + i
		}
	}
	return max+1 >= len(nums)
}

// @lc code=end
