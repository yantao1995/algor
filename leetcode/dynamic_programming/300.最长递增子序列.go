package leetcode

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxLength := 0
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if maxLength < dp[i] {
			maxLength = dp[i]
		}
	}
	return maxLength
}

// @lc code=end

/*
	dp记录到当前位置，的连续最大递增的个数
	依次从dp[i]向前寻找，如果当前值大于大于dp[j],
	那么就比较当前的最大值是否比dp[j]+1还大，因为要算上dp[i],所以+1。
	然后比较到最后，就是最大值。
	maxLength是为了保证，找到最长的一条递增链，因为不一定最后一个就是最大的。
*/
