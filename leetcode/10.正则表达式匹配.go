package leetcode

/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s))
	for k := range dp {
		dp[k] = make([]bool, len(p))
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if s[i] == p[j] || p[j] == '.' {
				dp[i][j] = true
			} else if p[j] == '*' {
				dp[i][j] = true
			}
		}
	}
	return dp[len(s)-1][len(p)-1]
}

// @lc code=end
