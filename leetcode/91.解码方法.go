package leetcode

/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([][]int, len(s))
	for k := range dp {
		dp[k] = make([]int, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '0' {
			dp[i][i] = 1
			for j := i + 1; j+1 < len(s); j++ {
				if s[j] == '0' {
					dp[i][j] = dp[i][j-1]
				} else if s[j-1] == '1' || (s[j-1] == '2' && s[j] < '7') {
					dp[i][j]++
					dp[i][len(dp)-1] = max(dp[i][len(dp)-1], dp[i][j]*dp[j+1][len(dp)-1])
				}
			}
		}
	}
	for k := range dp {
		for kk := range dp[k] {
			print(dp[k][kk])
		}
		println()
	}
	return dp[0][len(dp)-1]
}

// @lc code=end
