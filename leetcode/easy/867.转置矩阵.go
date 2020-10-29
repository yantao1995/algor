package leetcode

/*
 * @lc app=leetcode.cn id=867 lang=golang
 *
 * [867] 转置矩阵
 */

// @lc code=start
func transpose(A [][]int) [][]int {
	result := [][]int{}
	for j := 0; j < len(A[0]); j++ {
		temp := []int{}
		for i := 0; i < len(A); i++ {
			temp = append(temp, A[i][j])
		}
		result = append(result, temp)
	}
	return result
}

// @lc code=end
