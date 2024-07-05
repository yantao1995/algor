package lcpr

/*
 * @lc app=leetcode.cn id=3033 lang=golang
 * @lcpr version=30204
 *
 * [3033] 修改矩阵
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func modifiedMatrix(matrix [][]int) [][]int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxCol := map[int]int{}
	cols := map[int][]int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == -1 {
				cols[j] = append(cols[j], i)
			} else {
				maxCol[j] = max(maxCol[j], matrix[i][j])
			}
		}
	}
	for j, is := range cols {
		for _, i := range is {
			matrix[i][j] = maxCol[j]
		}
	}
	return matrix
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,-1],[4,-1,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[3,-1],[5,2]]\n
// @lcpr case=end

*/

/*
	记录坐标和列最大值即可
*/
