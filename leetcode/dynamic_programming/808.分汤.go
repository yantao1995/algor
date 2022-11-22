package leetcode

/*
 * @lc app=leetcode.cn id=808 lang=golang
 *
 * [808] 分汤
 */

// @lc code=start
func soupServings(n int) float64 {
	n = (n + 24) / 25
	if n >= 179 {
		return 1
	}
	dp := make([][]float64, n+1)
	for k := range dp {
		dp[k] = make([]float64, n+1)
		dp[0][k] = 1
	}
	dp[0][0] = 0.5
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = (dp[max(0, i-4)][j] + dp[max(0, i-3)][max(0, j-1)] +
				dp[max(0, i-2)][max(0, j-2)] + dp[max(0, i-1)][max(0, j-3)]) / 4
		}
	}
	return dp[n][n]
}

// @lc code=end

/*
	参考官方题解
	dp[i][j] 表示 A剩i，B剩j
	无法完成条件
		1. j==0,i>0
		2. i==0,j>0
	边界 i==0 ,j ==0时，a与b同时完成为1 ， 汤a先分配完为0， 总概率为 0.5


	状态转移： dp(i,j)= 0.25×(dp(i−4,y)+dp(i−3,y−1)+dp(i−2,y−2)+dp(i−1,y−3))
*/
