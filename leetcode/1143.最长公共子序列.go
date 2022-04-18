package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) < len(text2) {
		text1, text2 = text2, text1
	}
	text1 = " " + text1
	text2 = " " + text2
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
	for i := 1; i < len(text1); i++ {
		for j := 1; j < len(text2); j++ {
			dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			if text1[i] == text2[j] {
				dp[i][j] = max(dp[i][j], dp[i][j-1]+1)
			}
		}
	}
	fmt.Println(dp)
	return dp[len(dp)-1][len(dp[len(dp)-1])-1]
}

// @lc code=end
