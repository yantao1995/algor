package leetcode

/*
 * @lc app=leetcode.cn id=673 lang=golang
 *
 * [673] 最长递增子序列的个数
 */

// @lc code=start
func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxLength := 1
	dp, count := make([]int, len(nums)), make([]int, len(nums))
	dp[0], count[0] = 1, 1
	for i := 1; i < len(nums); i++ {
		dp[i], count[i] = 1, 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			}
		}
		if maxLength < dp[i] {
			maxLength = dp[i]
		}
	}
	result := 0
	for k := range dp {
		if dp[k] == maxLength {
			result += count[k]
		}
	}
	return result
}

// @lc code=end

/*
	dp[i]记录到i位置的连续最大递增的个数,count[i]记录当前nums[i]的最长递增个数
	nums[i] > nums[j]时,
	如果dp[j]+1 >dp[i]，那么当前dp[i]的最大值就是dp[j]+1,并且记录当前的count数量
	如果相等，那么就只累加个数
*/
