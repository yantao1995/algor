package leetcode

/*
 * @lc app=leetcode.cn id=1615 lang=golang
 *
 * [1615] 最大网络秩
 */

// @lc code=start
func maximalNetworkRank(n int, roads [][]int) int {
	matrix := make([][]int, n)
	for k := range matrix {
		matrix[k] = make([]int, n)
	}
	count := make([]int, n)
	for _, rd := range roads {
		matrix[rd[0]][rd[1]] = 1
		matrix[rd[1]][rd[0]] = 1
		count[rd[0]]++
		count[rd[1]]++
	}

	result := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i != j && count[i]+count[j]-matrix[i][j] > result {
				result = count[i] + count[j] - matrix[i][j]
			}
		}
	}
	return result
}

// @lc code=end

/*
	直接用矩阵连通枚举即可
*/
