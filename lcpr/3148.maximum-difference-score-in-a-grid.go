package lcpr

import "math"

/*
 * @lc app=leetcode.cn id=3148 lang=golang
 * @lcpr version=30204
 *
 * [3148] 矩阵中的最大得分
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
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
