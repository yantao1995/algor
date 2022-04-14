package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	sort.Ints(coins)
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] > i {
				break
			}
			dp[i] = dp[coins[j]-i] + 1
		}

	}
	return dp[amount]
}

// @lc code=end
