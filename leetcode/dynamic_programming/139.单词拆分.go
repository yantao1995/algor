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

/*
	初始化dp数据，即从第0个到每一位，能否用字典生成，能则dp[j]为true。
	开始遍历s，如果第i-1位可以用字典生成，而且i到j位，也可以用字典生成，
	那么第j位也应该为true，依次生成所有可能的j。
	于是dp最后一位，即为是否能够生成。
*/
