package leetcode

/*
 * @lc app=leetcode.cn id=2319 lang=golang
 *
 * [2319] 判断矩阵是否是一个 X 矩阵
 */

// @lc code=start
func checkXMatrix(grid [][]int) bool {
	n := len(grid) - 1
	for k := range grid {
		for kk := range grid {
			if k == kk || k+kk == n {
				if grid[k][kk] == 0 {
					return false
				}
			} else if grid[k][kk] != 0 {
				return false
			}
		}
	}
	return true
}

// @lc code=end

/*
	对角线就是 k==kk  和  k+kk == n
	直接判断即可
*/
