package leetcode

/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for k := range dp {
		dp[k] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i][j-1]
				}
			} else {
				if j == 0 {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j] + dp[i][j-1]
				}
			}
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end

/*
	二维dp数组记录当前存在的路径数量，初始化dp[0][0]为1
	因为机器人每次只能向下或者向右移动一步，所以当前
	dp[i][j] = dp[i-1][j] + dp[i][j-1]
	注意边界，适当修改  dp[i-1][j] + dp[i][j-1] 即可
*/
