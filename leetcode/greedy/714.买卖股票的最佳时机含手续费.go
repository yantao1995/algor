package leetcode

/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {
	dp := make([]int, len(prices))
	min, mini := prices[0], 0
	for i := 1; i < len(prices); i++ {
		dp[i] = dp[i-1]
		if dp[i] < dp[mini]+prices[i]-min-fee {
			dp[i] = dp[mini] + prices[i] - min - fee
			mini = i
			min = prices[i] - fee
		}
		if prices[i] < min {
			mini = i
			min = prices[i]
		}
	}
	return dp[len(dp)-1]
}

// @lc code=end

/*
	动归 + 贪心
	动归 是   由 最小值索引(mini) 买入，再由当前索引(i) 卖出，所能得到的累积最大值

	贪心 是  如果 最小值索引 (mini) 未被 更新，也就是当前 计算动归值时，
	还是用的上次的值计算，那么就不需要交手续费，所以在计算完dp[i]时，min = prices[i] - fee，
	表示下一轮未出手续费，因为 当前值减少了手续费的大小，所以下一轮出手续费时就抵扣了

	而遇到更小值 prices[i] 时，后续计算都是从这个 i开始的，所以需要给一轮手续费，即 min = prices[i]
*/
