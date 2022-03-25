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
		for j := 0; j < i; j++ { //j -> i 的当前最大值
			if prices[j] < prices[i] {
				if j >= 2 {
					dp[i] = max(dp[i], dp[j-2]+prices[i]-prices[j])
				} else {
					dp[i] = max(dp[i], prices[i]-prices[j])
				}
			}
		}
		dp[i] = max(dp[i], dp[i-1])
	}
	return dp[len(prices)-1]
}

// @lc code=end
