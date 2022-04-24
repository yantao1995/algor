package leetcode

/*
 * @lc app=leetcode.cn id=516 lang=golang
 *
 * [516] 最长回文子序列
 */

// @lc code=start
func longestPalindromeSubseq(s string) int {
	dp := make([][]bool, len(s)) //记录哪些是回文
	for k := range dp {
		dp[k] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i < 2 || dp[i+1][j-1]) {
				dp[i][j] = true
			}
		}
	}
	dp2 := make([][]int, len(s)) //记录回文的长度
	for k := range dp2 {
		dp2[k] = make([]int, len(s))
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if dp[i][j] {
				dp2[i][j] = j - i + 1
			} else if s[i] == s[j] {
				dp2[i][j] = dp2[i+1][j-1] + 2
			} else {
				dp2[i][j] = max(dp2[i][j-1], dp2[i+1][j])
			}
		}
	}
	return dp2[0][len(s)-1]
}

// @lc code=end

/*
	dp[i][j] 记录索引 i到j 是否为回文串  【做完之后发现dp可有可无，dp可以用来解前序(回文子串)题】
	dp2[i][j] 记录索引  i到j 的回文子串的最大长度

	例:
		a  b  c  b  a
	  a 1  1  1  3  5
	  b    1  1  3  3
	  c       1  1  1
	  b          1  1
	  a             1

	可以看出来，如果当前串是回文，那么长度就是  j-i+1
	如果不是回文串:
		1.  s[i] == s[j] ,那么可以保留s[i]和s[j],所以长度应该是 s[i+1]到s[j-1]的长度加上 s[i]与s[j] ，
			即：dp2[i][j] = dp2[i+1][j-1] + 2
		2.  s[i] != s[j] ,那么应该抛弃s[i]或者s[j]或者全抛,长度应该是 从s[i+1]到s[j] 获取 s[i]到s[j-1]的最大长度，
			又因为在计算 s[i+1]到s[j] 和 s[i]到s[j-1 已经抛过头和尾了，所以不需要单独取全抛
			即:dp2[i][j] = max(dp2[i][j-1], dp2[i+1][j])
*/
