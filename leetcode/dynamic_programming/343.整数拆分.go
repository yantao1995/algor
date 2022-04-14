package leetcode

/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start
func integerBreak(n int) int {
	dp := make([]int, n+2)
	dp[2], dp[3] = 1, 2

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 4; i <= n; i++ {
		for j := i - 2; j >= i/2; j-- {
			dp[i] = max(dp[i], max(dp[j], j)*max(dp[(i-j)], (i-j)))
		}
	}
	return dp[n]
}

// @lc code=end

/*
	先找规律
	n值		 2 3 4 5 6 7  8  9  10 11 12
	最大值	 1 2 4 6 9 12 18 27 36 54 81
	dp[i] 表示存当前 i 的组成的最大乘积

	比如 8 : 可以 是 4+4  5+3 6+2 ，然后在这些组合中取最大乘积

	比如遍历到 6+2的组合， 那么可以拆分为  6*dp[2] 或者 dp[6] * dp[2] 或者 dp[6]*2 或者 6 *2 ，
	只需要取这些组合中的最大值，那么即为  max(dp[6],6) * max(dp[2],2)

	然后依次向后取到目标值
*/
