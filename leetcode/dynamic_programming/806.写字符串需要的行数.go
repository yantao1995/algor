package leetcode

/*
 * @lc app=leetcode.cn id=806 lang=golang
 *
 * [806] 写字符串需要的行数
 */

// @lc code=start
func numberOfLines(widths []int, s string) []int {
	dp := []int{1, 0}
	bm := map[byte]int{}
	for k := range widths {
		bm['a'+uint8(k)] = widths[k]
	}
	for i := 0; i < len(s); i++ {
		dp[1] = bm[s[i]] + dp[1]
		if dp[1] > 100 {
			dp[0]++
			dp[1] = bm[s[i]]
		}
	}
	return dp
}

// @lc code=end

/*
	动态规划,dp[0]表示需要多少行，dp[1]表示当前行的占用长度
	如果当前行加上当前字符长度大于100，那么应该直接加到下一行
	即 dp[0]++ && dp[1] = 当前字符长度
*/
