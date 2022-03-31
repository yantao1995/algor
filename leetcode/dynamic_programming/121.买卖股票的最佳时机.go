package leetcode

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
// func maxProfit(prices []int) int {
// 	minPrice := 0
// 	max := 0
// 	if len(prices) < 1 {
// 		return max
// 	}
// 	minPrice = prices[0]
// 	for i := 1; i < len(prices); i++ {
// 		if prices[i] < minPrice {
// 			minPrice = prices[i]
// 		} else if prices[i]-minPrice > max {
// 			max = prices[i] - minPrice
// 		}
// 	}
// 	return max
// }
// @lc code=end

/*
	因为买了才能卖，只需要再最低买入，最高卖出即可，
	所以当前买入后，后面价格比现在高就可以卖，但是要和之前的历史最高价作比较，得到一个最高的差价
*/
