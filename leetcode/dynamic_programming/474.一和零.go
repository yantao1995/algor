package leetcode

/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for k := range dp {
		dp[k] = make([]int, n+1)
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	count0, count1 := 0, 0
	for k := range strs {
		count0, count1 = 0, 0
		for j := range strs[k] {
			if strs[k][j] == '1' {
				count1++
			} else {
				count0++
			}
		}
		for i := m; i >= count0; i-- { //从后往前，防止重复计算
			for j := n; j >= count1; j-- {
				dp[i][j] = max(dp[i][j], dp[i-count0][j-count1]+1)
			}
		}
	}
	return dp[m][n]
}

// @lc code=end

/*
	dp[i][j] 表示当前i个0  j个1 的子集数量
	dp max(当前的，-count0 和 -count1 的数量 + 当前子集)

	i,j 从后往前，防止 dp[i-count0][j-count1] 重复计算
*/
