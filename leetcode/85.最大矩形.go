package leetcode

/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	max := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				current := 1
				for broad := 1; i+broad < len(matrix) && j+broad < len(matrix[i+broad]); broad++ {
					for ii := i + broad; ii < len(matrix); ii++ {
						for jj := j + broad; jj < len(matrix[ii]); jj++ {
							if matrix[ii][jj] == '0' {
								goto lab
							}
						}
					}
					current++
					if current > max {
						max = current
					}
				}

			}
		lab:
		}
	}
	return max
}

// @lc code=end
