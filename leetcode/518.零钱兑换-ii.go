package leetcode

/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
func change(amount int, coins []int) int {
	// sort.Ints(coins)
	dp := make([][]int, amount+1) // dp[i][j]
	// for k := range dp {
	// 	dp[k] = make([]int, amount+1)
	// }
	// max := func(a, b int) int {
	// 	if a > b {
	// 		return a
	// 	}
	// 	return b
	// }
	// for k := range coins { //用当前硬币 有多少种

	// }

	return dp[amount][amount]
}

// @lc code=end
