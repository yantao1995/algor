package leetcode

import "fmt"

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

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp1 := make([]int, len(prices))
	dp2 := make([]int, len(prices))
	dp3 := make([]int, len(prices))
	dp4 := make([]int, len(prices))
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp1[0] = 0 - prices[0]
	for i := 1; i < len(prices); i++ {
		dp1[i] = max(dp1[i-1], 0-prices[i])
		dp2[i] = max(dp2[i-1], prices[i]+dp1[i])
		dp3[i] = max(dp3[i-1], dp2[i]-prices[i])
		dp4[i] = max(dp2[i], max(dp4[i-1], prices[i]+dp3[i]))
	}
	fmt.Println(dp1)
	fmt.Println(dp2)
	fmt.Println(dp3)
	fmt.Println(dp4)
	return dp4[len(dp4)-1]
}

// @lc code=end
