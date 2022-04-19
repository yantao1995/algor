package leetcode

import (
	"fmt"
	"sort"
)

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
	pft := make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		pft[i] = prices[i] - prices[i-1]
		if pft[i-1] > 0 && pft[i] > 0 {
			pft[i] += pft[i-1]
		}
	}
	fmt.Println(pft)
	sort.Ints(pft)
	return pft[len(pft)-1] + pft[len(pft)-2]
}

// @lc code=end
