package leetcode

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	memo := make([][]int, len(prices))
	for k := range memo {
		memo[k] = make([]int, len(prices))
	}
	var dp func(i int, isBuy, isSell bool)
	dp = func(i int, isBuy, isSell bool) {
		// if isSell {

		// }
	}
	dp(0, true, false)
	dp(0, false, false)
	if memo[len(prices)][len(prices)] < 0 {
		return 0
	}
	return memo[len(prices)][len(prices)]
}

// @lc code=end
