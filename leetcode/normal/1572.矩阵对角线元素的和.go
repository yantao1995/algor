package leetcode

/*
 * @lc app=leetcode.cn id=1572 lang=golang
 *
 * [1572] 矩阵对角线元素的和
 */

// @lc code=start
func diagonalSum(mat [][]int) int {
	squrt := len(mat)
	result := 0
	for k := range mat {
		result += mat[k][k]
		result += mat[k][squrt-k-1]
	}
	if squrt%2 == 1 {
		result -= mat[squrt/2][squrt/2]
	}
	return result
}

// @lc code=end
