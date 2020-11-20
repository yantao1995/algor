package leetcode

/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, S int) int {
	count := 0
	dp494(0, 0, S, nums, &count)
	return count
}

func dp494(total, index, S int, nums []int, count *int) {
	if index == len(nums) {
		if total == S {
			*count++
		}
		return
	}
	dp494(total-nums[index], index+1, S, nums, count)
	dp494(total+nums[index], index+1, S, nums, count)
}

// @lc code=end
