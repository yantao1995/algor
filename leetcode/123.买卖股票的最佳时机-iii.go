package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for k := range dp {
		dp[k] = make([]int, 2) //0第一次卖出  1第二次卖出
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			dp[j][0] = max(dp[j][0], dp[j-1][0]+prices[j]-prices[j-1])
		}
	}
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			dp[j][1] = max(dp[j][0], dp[j-1][0]+prices[j]-prices[j-1])
		}
	}
	fmt.Println(dp)
	return dp[len(dp)-1][1]
}

// @lc code=end
