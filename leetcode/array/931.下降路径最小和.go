package leetcode

/*
 * @lc app=leetcode.cn id=931 lang=golang
 *
 * [931] 下降路径最小和
 */

// @lc code=start
func minFallingPathSum(matrix [][]int) int {
	result, min := 0, 0
	for i := 0; i < len(matrix); i++ {
		result = matrix[i][0]
		if i > 0 {
			result += matrix[i-1][0]
		}
		for j := 0; j < len(matrix[i]); j++ {
			if i > 0 {
				min = matrix[i-1][j]
				if j-1 >= 0 && matrix[i-1][j-1] < min {
					min = matrix[i-1][j-1]
				}
				if j+1 < len(matrix[i]) && matrix[i-1][j+1] < min {
					min = matrix[i-1][j+1]
				}
				matrix[i][j] += min
			}
			if result > matrix[i][j] {
				result = matrix[i][j]
			}
		}
	}
	return result
}

// @lc code=end
