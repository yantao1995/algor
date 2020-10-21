package leetcode

/*
 * @lc app=leetcode.cn id=1380 lang=golang
 *
 * [1380] 矩阵中的幸运数
 */

// @lc code=start
func luckyNumbers(matrix [][]int) []int {
	result := []int{}
	for i := 0; i < len(matrix); i++ {
		temp := matrix[0][0]
		tempJ := 0
		for j := 0; j < len(matrix[i]); j++ {
			if temp > matrix[i][j] {
				temp = matrix[i][j]
				tempJ = j
			}
		}
		tempMax := matrix[0][0]
		tempI := 0
		for k := range matrix {
			if tempMax < matrix[k][tempJ] {
				tempMax = matrix[k][tempJ]
				tempI = k
			}
		}
		if tempI == i {
			result = append(result, matrix[tempI][tempJ])
		}
	}
	return result
}

// @lc code=end
