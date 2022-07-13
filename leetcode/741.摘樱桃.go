package leetcode

/*
 * @lc app=leetcode.cn id=741 lang=golang
 *
 * [741] 摘樱桃
 */

// @lc code=start
func cherryPickup(grid [][]int) int {
	n := len(grid)
	dp := [3][][]int{}
	for k := range dp {
		dp[k] = make([][]int, n)
		for kk := range dp[k] {
			dp[k][kk] = make([]int, n)
		}
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp[0][0][0] = grid[0][0]
	for step := 0; step < 2*n-1; step++ { //当前走的步数
		for num := 1; num < 3; num++ {
			for i := 0; i <= step && i < n; i++ {
				for j := step - i; j < n; j++ { //处于左下到右上的斜线
					if grid[i][j] != -1 {
						if i == 0 && j == 0 {
							dp[num][i][j] = grid[i][j]
						} else if i == 0 {
							dp[num][i][j] = max(dp[num-1][i][j], dp[num-1][i][j-1]+grid[i][j])
						} else if j == 0 {
							dp[num][i][j] = max(dp[num-1][i][j], dp[num-1][i-1][j]+grid[i][j])
						} else {
							dp[num][i][j] = max(dp[num-1][i][j], max(dp[num-1][i-1][j], dp[num-1][i][j-1])+grid[i][j])
						}
					}
				}
			}
		}
	}
	return dp[2][n-1][n-1]
}

// @lc code=end
