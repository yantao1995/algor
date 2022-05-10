package leetcode

/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	setZero := func(i, j int) {
		for ii := 0; ii < len(matrix); ii++ {
			matrix[ii][j] = 0
		}
		for jj := 0; jj < len(matrix[i]); jj++ {
			matrix[i][jj] = 0
		}
	}
	mi, mj := map[int]bool{}, map[int]bool{}
	for k := range matrix {
		for kk := range matrix[k] {
			if matrix[k][kk] == 0 {
				mi[k] = true
				mj[kk] = true
			}
		}
	}
	for i := range mi {
		for j := range mj {
			setZero(i, j)
		}
	}
}

// @lc code=end

/*
	mi，mj 分别记录哪些行，哪些列需要被置0
	然后分别对行列分别置0就行了
*/
