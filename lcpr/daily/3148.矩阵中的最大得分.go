package lcpr

/*
 * @lc app=leetcode.cn id=3148 lang=golang
 * @lcpr version=30204
 *
 * [3148] 矩阵中的最大得分
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// func maxScore(grid [][]int) int {
// 	maxMatrix := make([][]int, len(grid))
// 	for i := 0; i < len(grid); i++ {
// 		maxMatrix[i] = make([]int, len(grid[i]))
// 	}
// 	for i := len(grid) - 1; i >= 0; i-- {
// 		for j := len(grid[i]) - 1; j >= 0; j-- {
// 			if i < len(grid)-1 && j < len(grid[i])-1 {
// 				maxMatrix[i][j] = max(grid[i][j], maxMatrix[i+1][j], maxMatrix[i][j+1])
// 			} else if i < len(grid)-1 {
// 				maxMatrix[i][j] = max(grid[i][j], maxMatrix[i+1][j])
// 			} else if j < len(grid[i])-1 {
// 				maxMatrix[i][j] = max(grid[i][j], maxMatrix[i][j+1])
// 			} else {
// 				maxMatrix[i][j] = grid[i][j]
// 			}
// 		}
// 	}
// 	result := math.MinInt
// 	for i := 0; i < len(grid); i++ {
// 		for j := 0; j < len(grid[i]); j++ {
// 			if i < len(grid)-1 {
// 				result = max(result, maxMatrix[i+1][j]-grid[i][j])
// 			}
// 			if j < len(grid[i])-1 {
// 				result = max(result, maxMatrix[i][j+1]-grid[i][j])
// 			}
// 		}
// 	}
// 	return result
// }

// @lc code=end

/*
// @lcpr case=start
// [[9,5,7,3],[8,9,6,1],[6,7,14,3],[2,5,3,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[4,3,2],[3,2,1]]\n
// @lcpr case=end

*/

/*
从右下角到左上角遍历，依次记录从右下角到当前值的最大值。然后遍历计算即可



直接挨个取数比较，超时
Time Limit Exceeded
556/564 cases passed (N/A)

func maxScore(grid [][]int) int {
	result := math.MinInt
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {

			for m := i; m < len(grid); m++ {
				for n := j; n < len(grid[m]); n++ {
					if i != m || j != n {
						result = max(result, grid[m][n]-grid[i][j])
					}
				}
			}
		}
	}
	return result
}
*/
