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
	for i := 0; i < len(text2); i++ {
		if text1[0] == text2[i] {
			dp[0][i] = 1
		}
		if i > 0 {
			dp[0][i] = max(dp[0][i], dp[0][i-1])
		}
	}
	for i := 0; i < len(text1); i++ {
		if text2[0] == text1[i] {
			dp[i][0] = 1
		}
		if i > 0 {
			dp[i][0] = max(dp[i][0], dp[i-1][0])
		}
	}

	for i := 1; i < len(text1); i++ {
		for j := 1; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[len(text1)-1][len(text2)-1]
}

// @lc code=end

/*
	dp[i][j] 表示到text1[i] 到text2[j] 的最长公共子序列
	如果当前字符相等，那么直接累加 dp[i-1][j-1],否则 i-1与j-1取最大值
*/
