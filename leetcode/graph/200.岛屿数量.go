package leetcode

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	dp := make([][]bool, len(grid))
	count := 0
	for k := range grid {
		dp[k] = make([]bool, len(grid[k]))
	}
	var dfs func(i, j int)
	dfs = func(i, j int) { //右下标记岛屿
		if i < 0 || i >= len(grid) ||
			j < 0 || j >= len(grid[i]) ||
			grid[i][j] != '1' || dp[i][j] {
			return
		}
		dp[i][j] = true
		dfs(i+1, j)
		dfs(i, j+1)
		dfs(i-1, j)
		dfs(i, j-1)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && !dp[i][j] {
				dfs(i, j)
				count++
			}
		}
	}
	return count
}

// @lc code=end
