package easy

/*
 * @lc app=leetcode.cn id=1380 lang=golang
 *
 * [1380] 矩阵中的幸运数
 */

// @lc code=start
func luckyNumbers(matrix [][]int) []int {
	result := []int{}
	for i := 0; i < len(matrix); i++ {
		temp := matrix[i][0]
		tempJ := 0
		for j := 0; j < len(matrix[i]); j++ {
			if temp > matrix[i][j] {
				temp = matrix[i][j]
				tempJ = j
			}
		}
		tempMax := matrix[0][tempJ]
		tempI := 0
		for ii := 0; ii < len(matrix); ii++ {
			if tempMax < matrix[ii][tempJ] {
				tempMax = matrix[ii][tempJ]
				tempI = ii
			}
		}
		if tempI == i {
			result = append(result, matrix[tempI][tempJ])
		}
	}
	return result
}

// @lc code=end
