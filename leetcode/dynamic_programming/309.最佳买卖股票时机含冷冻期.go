package leetcode

/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	if len(prices) < 2 {
		return 0
	}
	//从 第0个 到第 i个。当前能取到的最大值
	for i := 1; i < len(prices); i++ {
		lessIndex := []int{}
		for j := 0; j < i; j++ { //j -> i 的当前最大值
			if prices[j] < prices[i] {
				lessIndex = append(lessIndex, j)
			}
		}
		for _, index := range lessIndex {
			if index >= 2 {
				dp[i] = max(dp[i], dp[index-2]+prices[i]-prices[index])
			} else {
				dp[i] = max(dp[i], prices[i]-prices[index])
			}
		}
		dp[i] = max(dp[i], dp[i-1])
	}
	return dp[len(prices)-1]
}

// @lc code=end
