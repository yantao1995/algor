package leetcode

/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	max := 0
	check := func(i, j, board int) bool {
		for ii := i; ii <= i+board; ii++ {
			if matrix[ii][j+board] == '0' {
				return false
			}
		}
		for jj := j; jj <= j+board; jj++ {
			if matrix[i+board][jj] == '0' {
				return false
			}
		}
		return true
	}
	for i := 0; i < len(matrix)-max; i++ {
		for j := 0; j < len(matrix[i])-max; j++ {
			if matrix[i][j] == '1' {
				current := 1
				for i+current < len(matrix) && j+current < len(matrix[i]) && check(i, j, current) {
					current++
				}
				if current > max {
					max = current
				}
			}
		}
	}
	return max * max
}

// @lc code=end
