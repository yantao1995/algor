package leetcode

/*
 * @lc app=leetcode.cn id=940 lang=golang
 *
 * [940] 不同的子序列 II
 */

// @lc code=start
func distinctSubseqII(s string) int {
	mod := int(1e9 + 7)
	dp := [26]int{}
	for k := range s {
		total := 1
		for _, v := range dp {
			total = (total + v) % mod
		}
		dp[s[k]-'a'] = total
	}
	for i := 1; i < len(dp); i++ {
		dp[0] = (dp[0] + dp[i]) % mod
	}
	return dp[0]
}

// @lc code=end

/*
	参考官方题解

	以当前字符作为所有子序列的末尾字符，可以构造出新的n个子序列，
	前提是前一个字符不同，如果前一个字符相同，那就前前个，依次向前推

*/
