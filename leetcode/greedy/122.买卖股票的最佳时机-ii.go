package leetcode

/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
// func maxProfit(prices []int) int {
// 	result := 0
// 	for i := 1; i < len(prices); i++ {
// 		if prices[i] >= prices[0] {
// 			result += prices[i] - prices[0]
// 		}
// 		prices[0] = prices[i]
// 	}
// 	return result
// }

// @lc code=end

/*
	贪心
	因为可以同一天出售，所以只要后一天比当天高，就赚差价
	用 price[0]当哨兵，每次把当天的价格存到price[0]里
*/
