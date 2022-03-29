package leetcode

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	dictMap := map[string]bool{}
	for k := range wordDict {
		dictMap[wordDict[k]] = true
	}
	dp := make([]bool, len(s))
	for j := 1; j <= len(s); j++ {
		if dictMap[s[0:j]] {
			dp[j-1] = true
		}
	}
	for i := 1; i < len(s); i++ {
		for j := i + 1; j <= len(s) && dp[i-1]; j++ {
			if dictMap[s[i:j]] {
				dp[j-1] = true
			}
		}
	}
	return dp[len(dp)-1]
}

// @lc code=end
