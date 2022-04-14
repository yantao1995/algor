package leetcode

/*
 * @lc app=leetcode.cn id=1049 lang=golang
 *
 * [1049] 最后一块石头的重量 II
 */

// @lc code=start
func lastStoneWeightII(stones []int) int {
	total := 0
	for k := range stones {
		total += stones[k]
	}
	half := total / 2
	dp := make([]bool, total+1)
	for i := 0; i < len(stones); i++ {
		for j := len(dp) - 1; j > 0; j-- {
			if dp[j] && j+stones[i] < len(dp) {
				dp[j+stones[i]] = true
			}
		}
		dp[stones[i]] = true
	}
	for ; half > 0; half-- {
		if dp[half] && dp[total-half] {
			return total - half - half
		}
	}
	return total
}

// @lc code=end

/*
	和分割等和子集一个思路

	分别填充 [1-total] 的 dp 数组，表示当前dp[i]可以组成，即 i的值可以加出来

	最后分成两半来粉碎，两边值越接近，粉碎的值就越小

	那么可以取 half ，从中间来向两边拆开成两份
*/
