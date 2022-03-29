package leetcode

/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 {
				if j > 0 {
					grid[i][j] += grid[i][j-1]
				}
			} else {
				if j == 0 {
					grid[i][j] += grid[i-1][j]
				} else {
					grid[i][j] += min(grid[i-1][j], grid[i][j-1])
				}
			}
		}
	}
	return grid[len(grid)-1][len(grid[len(grid)-1])-1]
}

// @lc code=end
