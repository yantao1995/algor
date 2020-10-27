package easy

/*
 * @lc app=leetcode.cn id=1351 lang=golang
 *
 * [1351] 统计有序矩阵中的负数
 */

// @lc code=start
func countNegatives(grid [][]int) int {
	count := 0
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j] >= 0 {
				break
			} else {
				count++
			}
		}
	}
	return count
}

// @lc code=end
