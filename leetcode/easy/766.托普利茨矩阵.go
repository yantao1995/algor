package easy

/*
 * @lc app=leetcode.cn id=766 lang=golang
 *
 * [766] 托普利茨矩阵
 */

// @lc code=start
func isToeplitzMatrix(matrix [][]int) bool {
	w, l := len(matrix), len(matrix[0])
	for i := 0; i < w; i++ { //左下半区
		tempI := i
		j := 0
		for tempI+1 < w && j+1 < l {
			if matrix[tempI][j] != matrix[tempI+1][j+1] {
				return false
			}
			tempI++
			j++
		}
	}
	for j := 0; j < l; j++ { //右上半区
		tempJ := j
		i := 0
		for tempJ+1 < l && i+1 < w {
			if matrix[i][tempJ] != matrix[i+1][tempJ+1] {
				return false
			}
			tempJ++
			i++
		}
	}
	return true
}

// @lc code=end
