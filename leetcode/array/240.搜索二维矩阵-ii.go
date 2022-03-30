package leetcode

/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	ds, de := len(matrix), len(matrix[0])
	getNextSide := func(i, j, ds, de int) (int, int) {
		if i+1 < ds {
			i++
		}
		if j+1 < de {
			j++
		}
		return i, j
	}
	var scan func(i, j, ds, de int) bool
	scan = func(i, j, ds, de int) bool {
		tempI, tempJ := i, j
		for i < ds || j < de {
			if target == matrix[i][j] {
				return true
			} else if target > matrix[i][j] {
				if ii, jj := getNextSide(i, j, ds, de); ii == i && jj == j {
					return false
				} else {
					i, j = ii, jj
				}
			} else {
				if tempI == i && tempJ == j {
					return false
				}
				if scan(i, tempJ, ds, j) || scan(tempI, j, i, de) { //对角线斜边搜索，分块搜索
					return true
				}
				return false
			}
		}
		return false
	}
	return scan(0, 0, ds, de)
}

// @lc code=end
