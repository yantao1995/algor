package leetcode

/*
 * @lc app=leetcode.cn id=1605 lang=golang
 *
 * [1605] 给定行和列的和求可行矩阵
 */

// @lc code=start
func restoreMatrix(rowSum []int, colSum []int) [][]int {
	result := make([][]int, len(rowSum))
	for k := range result {
		result[k] = make([]int, len(colSum))
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			result[i][j] = min(rowSum[i], colSum[j])
			rowSum[i] -= result[i][j]
			colSum[j] -= result[i][j]
		}
	}
	return result
}

// @lc code=end

/*
	贪心算法，直接贪当前行列就是行与列的当前最小的一个值。这样不足的可以由其他行或列补齐
*/
