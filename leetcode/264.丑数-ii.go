package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	dp := []int{1}
	m := map[int]bool{}
	for i := 0; i <= n; i++ {
		if !m[dp[i]*2] {
			dp = append(dp, dp[i]*2)
			m[dp[i]*2] = true
		}
		if !m[dp[i]*3] {
			dp = append(dp, dp[i]*3)
			m[dp[i]*3] = true
		}
		if !m[dp[i]*5] {
			dp = append(dp, dp[i]*5)
			m[dp[i]*5] = true
		}
		sort.Ints(dp)
	}
	return dp[n-1]
}

// @lc code=end
