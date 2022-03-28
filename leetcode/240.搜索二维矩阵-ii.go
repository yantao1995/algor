package leetcode

/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	ds, de := 0, min(len(matrix), len(matrix[0]))
	dm := (ds + de) / 2
	for dm < len(matrix) && dm < len(matrix[0]) && ds < dm && dm < de {
		if matrix[dm][dm] == target {
			return true
		} else if matrix[dm][dm] < target {
			ds = dm
			dm = (dm + de) / 2
		} else {
			de = dm
			dm = (ds + dm) / 2
		}
	}

	return false
}

// @lc code=end
