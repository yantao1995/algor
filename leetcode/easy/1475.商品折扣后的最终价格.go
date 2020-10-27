package easy

/*
 * @lc app=leetcode.cn id=1475 lang=golang
 *
 * [1475] 商品折扣后的最终价格
 */

// @lc code=start
func finalPrices(prices []int) []int {
	for k := range prices {
		for i := k + 1; i < len(prices); i++ {
			if prices[i] <= prices[k] {
				prices[k] -= prices[i]
				break
			}
		}
	}
	return prices
}

// @lc code=end
