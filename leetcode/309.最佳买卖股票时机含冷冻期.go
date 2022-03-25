package leetcode

/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {
	max := 0
	if len(prices) < 1 {
		return 0
	}
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if max < prices[i]-minPrice {
			max = prices[i] - minPrice
		}
	}
	return max
}

// @lc code=end
