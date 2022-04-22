package leetcode

/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1))
	for k := range dp {
		dp[k] = make([]int, len(text2))
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i := 0; i < len(text2); i++ {
		if text1[0] == text2[i] {
			dp[0][i] = 1
		}
		if i > 0 {
			dp[0][i] = max(dp[0][i], dp[0][i-1])
		}
	}
	for i := 1; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				if j > 0 && i > 0 {
					dp[i][j] = max(min(dp[i-1][j], dp[i][j-1])+1, dp[i][j])
				} else {
					dp[i][j] = 1
				}
			} else {
				if j > 0 {
					dp[i][j] = max(dp[i][j-1], dp[i-1][j])
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	for k := range dp {
		for kk := range dp[k] {
			print(dp[k][kk], "\t")
		}
		println()
	}
	return dp[len(text1)-1][len(text2)-1]
}

// @lc code=end
