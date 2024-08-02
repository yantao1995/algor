package lcpr

/*
 * @lc app=leetcode.cn id=3128 lang=golang
 * @lcpr version=30204
 *
 * [3128] 直角三角形
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numberOfRightTriangles(grid [][]int) int64 {
	result := 0
	if len(grid) == 0 {
		return 0
	}
	rows, cols := make([]int, len(grid)), make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				rows[i]++
				cols[j]++
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				result += (rows[i] - 1) * (cols[j] - 1)
			}
		}
	}
	return int64(result)
}

// @lc code=end

/*
// @lcpr case=start
// [[0,1,0],[0,1,1],[0,1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,0,0,0],[0,1,0,1],[1,0,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,0,1],[1,0,0],[1,0,0]]\n
// @lcpr case=end

*/

/*
优化思路，提前数好个数


以当前点位直角。数上下左右分别有1的个数，然后计算

Time Limit Exceeded
585/637 cases passed (N/A)

func numberOfRightTriangles(grid [][]int) int64 {
	result := 0
	up, down, left, right := 0, 0, 0, 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				up, down, left, right = 0, 0, 0, 0
				for ti := i - 1; ti >= 0; ti-- {
					if grid[ti][j] == 1 {
						up++
					}
				}
				for ti := i + 1; ti < len(grid); ti++ {
					if grid[ti][j] == 1 {
						down++
					}
				}
				for tj := j - 1; tj >= 0; tj-- {
					if grid[i][tj] == 1 {
						left++
					}
				}
				for tj := j + 1; tj < len(grid[i]); tj++ {
					if grid[i][tj] == 1 {
						right++
					}
				}
				result += (up + down) * (left + right)
			}
		}
	}
	return int64(result)
}
*/
