package leetcode

/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
// func maxProfit(prices []int) int {   //// dp 超时
// 	if len(prices) < 2 {
// 		return 0
// 	}
// 	dp := make([]int, len(prices))
// 	max := func(a, b int) int {
// 		if a > b {
// 			return a
// 		}
// 		return b
// 	}
// 	minIndex, mProfit := 0, 0
// 	for i := 1; i < len(prices); i++ {
// 		if prices[i] < prices[minIndex] {
// 			minIndex = i
// 			dp[i] = dp[i-1]
// 		} else {
// 			dp[i] = max(dp[i-1], prices[i]-prices[minIndex])
// 		}
// 	}
// 	for i := 0; i < len(prices); i++ {
// 		for j := i + 1; j < len(prices); j++ {
// 			if dp[i]+prices[j]-prices[i] > mProfit {
// 				mProfit = dp[i] + prices[j] - prices[i]
// 			}
// 		}
// 	}

// 	fmt.Println(dp)
// 	return mProfit
// }

/////

// func maxProfit(prices []int) int {   // 时间复杂度降为 n  不超时
// 	if len(prices) < 2 {
// 		return 0
// 	}
// 	dp1 := make([]int, len(prices))
// 	dp2 := make([]int, len(prices))
// 	dp3 := make([]int, len(prices))
// 	dp4 := make([]int, len(prices))
// 	max := func(a, b int) int {
// 		if a > b {
// 			return a
// 		}
// 		return b
// 	}
// 	for k := range dp1 {
// 		dp1[k] = -100000
// 		dp3[k] = -100000
// 	}
// 	for i := 0; i < len(prices); i++ {
// 		if i == 0 {
// 			dp1[i] = 0 - prices[i]
// 			dp3[i] = 0 - prices[i]
// 		} else {
// 			dp1[i] = max(dp1[i-1], 0-prices[i])
// 			dp2[i] = max(dp2[i-1], prices[i]+dp1[i-1])
// 			dp3[i] = max(dp3[i-1], dp2[i-1]-prices[i])
// 			dp4[i] = max(dp4[i-1], prices[i]+dp3[i-1])
// 		}
// 	}
// 	return dp4[len(dp4)-1]
// }

// @lc code=end

/*
	dp[i] 表示当天的最大收益
	dp1 记录当天第一次买
	dp2 记录当天第一次卖
	dp3 记录当天第二次买
	dp4 记录当天第二次卖

	所以 dp2 依赖 dp1 ， dp3 依赖dp2  , dp4 依赖dp3
*/
