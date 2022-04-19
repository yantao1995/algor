package leetcode

/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
// func maxProfit(k int, prices []int) int {
// 	if len(prices) < 2 {
// 		return 0
// 	}
// 	dp := make([][]int, len(prices))
// 	for i := range dp {
// 		dp[i] = make([]int, k+1)
// 	}
// 	max := func(a, b int) int {
// 		if a > b {
// 			return a
// 		}
// 		return b
// 	}
// 	for i := 0; i < len(prices); i++ {
// 		for j := i + 1; j < len(prices); j++ {
// 			for m := 1; m <= k; m++ {
// 				if j > 0 {
// 					dp[j][m] = max(dp[j][m], dp[j-1][m])
// 				}
// 				dp[j][m] = max(dp[j][m], dp[i][m-1]+prices[j]-prices[i])
// 			}
// 		}
// 	}
// 	return dp[len(prices)-1][k]
// }

// @lc code=end

/*
	dp[i][k]   到第i天卖 k次的利润 //1第一次 2第二次 k次 只能从 k-1上累加
	每次的出售只能从上一轮的利润中加，就是代表当前是下一次,这样就限制了次数
*/
